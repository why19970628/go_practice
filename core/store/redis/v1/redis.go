package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/bluele/gcache"
	"github.com/go-kratos/aegis/topk"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	NoExpiration = time.Duration(0)
	Nil          = redis.Nil
)

// RedisOption 定义 Redis 配置选项
type (
	RedisOption func(*Redis) error

	Redis struct {
		Client        *redis.Client
		topK          topk.Topk
		localCache    gcache.Cache
		whiteList     map[string]struct{}
		localCacheTTL time.Duration
	}
)

func NewRedis(config RedisConfig, options ...RedisOption) (*Redis, error) {
	// 应用选项
	r := &Redis{}
	for _, option := range options {
		if err := option(r); err != nil {
			return nil, err
		}
	}
	if err := config.filDefault(); err != nil {
		return nil, err
	}
	r.Client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  config.DialTimeout,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		IdleTimeout:  config.IdleTimeout,
		MaxRetries:   config.MaxRetries,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		MaxConnAge:   config.MaxConnAge,
	})
	if r.localCacheTTL == 0 {
		r.localCacheTTL = config.LocalCacheTTL
	}

	if err := check(r); err != nil {
		return nil, err
	}
	return r, nil
}

func check(r *Redis) error {
	if r == nil {
		return errors.New("redis == nil")
	}
	if _, err := r.Client.Ping(context.Background()).Result(); err != nil {
		return err
	}
	return nil
}

// WithTopK 设置 TopK 选项
func WithTopK(topK topk.Topk) RedisOption {
	return func(r *Redis) error {
		r.topK = topK
		return nil
	}
}

// WithFirstLevelCache 设置一级缓存选项
func WithFirstLevelCache(cache gcache.Cache) RedisOption {
	return func(r *Redis) error {
		r.localCache = cache
		return nil
	}
}

// WithWhiteList 设置白名单选项
func WithWhiteList(whiteList map[string]struct{}) RedisOption {
	return func(r *Redis) error {
		r.whiteList = whiteList
		return nil
	}
}

// WithLocalCacheTTL 本地缓存过期时间
func WithLocalCacheTTL(localCacheTTL time.Duration) RedisOption {
	return func(r *Redis) error {
		r.localCacheTTL = localCacheTTL
		return nil
	}
}
