package redis

import (
	"github.com/bluele/gcache"
	"github.com/go-kratos/aegis/topk"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	rds *Redis
)

func initNewRedis() error {
	var err error
	rds, err = NewRedis(RedisConfig{
		Host:      "127.0.0.2:999",
		Port:      0,
		Password:  "",
		KeyPrefix: "test",
	},
		WithTopK(topk.NewHeavyKeeper(10, 10000, 5, 0.925, 0)),
		WithFirstLevelCache(gcache.New(100).LRU().Build()),
		WithWhiteList(map[string]struct{}{"123": {}}),
	)

	return err
}

func TestNewRedis(t *testing.T) {
	assert.Equal(t, nil, initNewRedis())
}
