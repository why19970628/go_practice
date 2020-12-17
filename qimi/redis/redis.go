package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ffmt.v1"
	"log"
	"strings"
	"sync"
	"time"
)

//var Client *redis.Client

func init() {
	//dbRedis := New()
	//Client = dbRedis.Client()
}

func New() *Redis {
	log.Print("redis.New Connect To Redis ...")
	// 链接上redis
	client := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolSize:    100,
		PoolTimeout: 500,
	})
	cmd := client.Ping()
	if cmd.Err() != nil {
		log.Fatal("redis.New client.Ping() Failed: %v", cmd.Err())
	}
	ticker := time.NewTicker(50)
	return &Redis{
		client:  client,
		Scripts: map[string]*redisScript{},
		ticker:  ticker,
	}
}

type redisScript struct {
	sync.Mutex
	tag  string
	src  string
	sha1 string
	err  error
}

type Redis struct {
	sync.RWMutex
	client  *redis.Client
	ticker  *time.Ticker
	Scripts map[string]*redisScript
}

func (rds *Redis) Client() *redis.Client {
	return rds.client
}


// 应用层应在初始化阶段完成所有LoadScript操作, 初始化后不应再调用此方法
func (rds *Redis) LoadScript(tag string, src string) error {
	rds.Lock()
	defer rds.Unlock()
	client := rds.Client()
	sha1, err := client.ScriptLoad(src).Result()
	if err != nil {
		log.Fatal("redis ScriptLoad '%s' failed: %v", tag, err)
	}
	rds.Scripts[tag] = &redisScript{
		tag:  tag,
		src:  src,
		sha1: sha1,
		err:  err,
	}
	return err
}


// 应用层应在初始化阶段完成所有LoadScript操作, 初始化后不应再调用此方法
func (rds *Redis) GetScript(tag string)  {
	ffmt.Puts(rds.Scripts[tag])
}

func (rds *Redis) getScriptSha1(tag string) (string, error) {
	rds.RLock()
	defer rds.RUnlock()
	if script, ok := rds.Scripts[tag]; ok {
		return script.sha1, script.err
	}
	return "", errors.New(fmt.Sprintf("invalid redis lua script tag: %s", tag))
}

func (rds *Redis) reLoadScript(tag string) error {
	rds.Lock()
	defer rds.Unlock()
	var (
		err  error = errors.New("no such script")
		sha1       = ""
	)
	if script, ok := rds.Scripts[tag]; ok {
		cmd := rds.Client().ScriptLoad(script.src)
		sha1, err = cmd.Result()
		script.sha1, script.err = sha1, err
	}

	if err != nil {
		log.Fatal("redis ReLoadScript '%s' failed: %v", tag, err)
	}
	return err
}

func (rds *Redis) EvalSha(tag string, keys []string, args ...interface{}) (interface{}, error) {
	var ret interface{} = nil
	sha1, err := rds.getScriptSha1(tag)
	if err == nil {
		cmd := rds.Client().EvalSha(sha1, keys, args...)
		ret, err = cmd.Result()
		if err != nil && strings.HasPrefix(err.Error(), "NOSCRIPT") {
			if err = rds.reLoadScript(tag); err == nil {
				if sha1, err = rds.getScriptSha1(tag); err == nil {
					cmd = rds.Client().EvalSha(sha1, keys, args...)
					ret, err = cmd.Result()
				}
			}
		}
	}

	return ret, err
}

func (rds *Redis) Eval(src string, keys []string, args ...interface{}) (interface{}, error) {
	cmd := rds.Client().Eval(src, keys, args...)
	ret, err := cmd.Result()
	return ret, err
}

func (rds *Redis) Close() error {
	return rds.client.Close()
}
