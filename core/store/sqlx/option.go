package sqlx

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

// Config options
type Config struct {
	Name string
	// DSN地址: mysql://root:secret@tcp(127.0.0.1:3307)/mysql?timeout=20s&readTimeout=20s
	DSN string `json:"dsn" toml:"dsn"`
	// Debug开关
	Debug bool `json:"debug" toml:"debug"`
	// 最大空闲连接数
	MaxIdleConns int `json:"maxIdleConns" toml:"maxIdleConns"`
	// 最大活动连接数
	MaxOpenConns int `json:"maxOpenConns" toml:"maxOpenConns"`
	// 连接的最大存活时间
	ConnMaxLifetime time.Duration `json:"connMaxLifetime" toml:"connMaxLifetime"`
	// 创建连接的错误级别，=panic时，如果创建失败，立即panic
	OnDialError string `json:"level" toml:"level"`
	// 慢日志阈值
	SlowThreshold time.Duration `json:"slowThreshold" toml:"slowThreshold"`
	// 拨超时时间
	DialTimeout time.Duration `json:"dialTimeout" toml:"dialTimeout"`
	// 自动使用影子表
	AutoShadowTable bool `json:"autoShadowTable" toml:"autoShadowTable"`
	raw             interface{}
	// 记录错误sql时,是否打印包含参数的完整sql语句
	// select * from aid = ?;
	// select * from aid = 288016;
	DetailSQL bool `json:"detailSql" toml:"detailSql"`
	// 重试次数
	Retry int `json:"retry" toml:"retry"`
	// 重试等待时间
	RetryWaitTime time.Duration `json:"retryWaitTime" toml:"retryWaitTime"`

	gormConfig gorm.Config `json:"-" toml:"-"`
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		DSN:             "",
		Debug:           false,
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: cast.ToDuration("300s"),
		OnDialError:     "panic",
		SlowThreshold:   cast.ToDuration("500ms"),
		DialTimeout:     cast.ToDuration("3s"),
		AutoShadowTable: false,
		raw:             nil,
		Retry:           2,
		RetryWaitTime:   cast.ToDuration("200ms"),
	}
}
