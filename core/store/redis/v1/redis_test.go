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
		Host:     "r-2zemnsr8cmr6trxnt6.redis.rds.aliyuncs.com",
		Port:     6379,
		Password: "qRf^Zzc3sBVzZcz2",
		//KeyPrefix: "test",
	},
		WithTopK(topk.NewHeavyKeeper(10, 10000, 5, 0.925, 0)),
		WithFirstLevelCache(gcache.New(100).LRU().Build()),
		WithWhiteList(map[string]interface{}{"test_white_key": "test_white_value"}),
	)

	return err
}

func TestNewRedis(t *testing.T) {
	assert.Equal(t, nil, initNewRedis())
}
