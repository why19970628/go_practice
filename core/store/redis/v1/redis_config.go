package redis

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultRedisPort     = 6379 // Default redis port configuration if not passed.
	DefaultDialTimeout   = time.Second * 10
	DefaultMaxRetries    = 3
	DefaultPoolSize      = 100
	DefaultReadTimeout   = time.Second * 3
	DefaultWriteTimeout  = time.Second * 3
	DefaultMaxConnAge    = time.Second * 60
	DefaultMinIdleConns  = 200
	DefaultLocalCacheTTL = time.Second * 60
)

type RedisConfig struct {
	Host          string        `yaml:"host" json:"host"`
	Port          int           `yaml:"port" json:"port"`
	Password      string        `yaml:"password" json:"password"`
	DB            int           `yaml:"db" json:"db"`
	KeyPrefix     string        `yaml:"key_prefix" json:"key_prefix"`
	MinIdleConns  int           `yaml:"min_idle_conns" json:"min_idle_conns"` // 最少空闲连接数
	PoolSize      int           `yaml:"pool_size" json:"pool_size"`           // 连接池上线
	IdleTimeout   time.Duration `yaml:"idle_timeout" json:"idle_timeout"`     // 一个连接多久未被连接认为该连接是空闲连接 单位：秒
	DialTimeout   time.Duration `yaml:"dial_timeout" json:"dial_timeout"`     //  Dial connection timeout.
	MaxRetries    int           `yaml:"max_retries" json:"max_retries"`       // 最大重试次数
	ReadTimeout   time.Duration `yaml:"read_timeout" json:"read_timeout"`     // 读超时时间 -1: 没有超时时间 0: 采用默认
	WriteTimeout  time.Duration `yaml:"write_timeout" json:"write_timeout"`   // 写超时时间 -1: 没有超时时间 0: 使用默认
	MaxConnAge    time.Duration `yaml:"max_conn_age" json:"max_conn_age"`     // Close connections older than this duration. If the value is zero, then the pool does not close connections based on age.
	LocalCacheTTL time.Duration `yaml:"local_cache_ttl" json:"local_cache_ttl"`
}

// 填充全部配置的默认值
func (c *RedisConfig) filDefault() error {
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == 0 {
		c.Port = DefaultRedisPort
	}
	if c.DialTimeout == 0 {
		c.DialTimeout = DefaultDialTimeout
	}
	if c.MaxRetries == 0 {
		c.MaxRetries = DefaultMaxRetries
	}
	if c.PoolSize == 0 {
		c.PoolSize = DefaultPoolSize
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = DefaultWriteTimeout
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = DefaultReadTimeout
	}
	if c.MaxConnAge == 0 {
		c.MaxConnAge = DefaultMaxConnAge
	}
	if c.MinIdleConns == 0 {
		c.MinIdleConns = DefaultMinIdleConns
	}
	if c.LocalCacheTTL == 0 {
		c.LocalCacheTTL = DefaultLocalCacheTTL
	}
	return c.resetDefault()
}

// 填充最小配置的默认值
func (c *RedisConfig) resetDefault() error {
	var isMatch bool
	var host string
	var port int
	var err error
	isMatch, host, port, err = formatAddr(c.Host)
	if err != nil {
		return err
	}
	if isMatch {
		c.Host = host
		c.Port = port
	}
	return nil
}

func formatAddr(s string) (match bool, addr string, port int, err error) {
	expr := ":(\\d+)$"
	var r *regexp.Regexp
	r, err = regexp.Compile(expr)
	if err != nil {
		return
	}
	ports := strings.TrimLeft(r.FindString(s), ":")
	if ports != "" {
		match = true
		iPort, eRrr := strconv.Atoi(ports)
		if eRrr != nil {
			return
		}
		addr = strings.TrimSuffix(s, fmt.Sprintf(":%s", ports))
		port = iPort
	}
	return
}
