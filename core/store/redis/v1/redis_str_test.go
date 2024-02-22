package redis

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisStr(t *testing.T) {
	err := initNewRedis()
	if err != nil {
		return
	}
	key1 := "key1"
	key2 := "key2"
	var (
		v1 = "val1"
		v2 = 456
	)
	fmt.Println(rds.GetTopKey())
	//assert.Equal(t, 1, len(rds.GetTopKey()))
	assert.Nil(t, rds.Set(context.Background(), key1, v1, time.Second*10))
	assert.Nil(t, rds.Set(context.Background(), key2, v2, time.Second*10))

	v1Res, _ := rds.Get(context.Background(), key1)
	assert.Equal(t, v1, v1Res)

	vRes, _ := rds.MGet(context.Background(), key1, key2)
	assert.Equal(t, 2, len(vRes))

	//delCount, err := rds.Del(context.Background(), key1, key2)
	//assert.Nil(t, err)
	//assert.Equal(t, int64(2), delCount)
}

func TestRedisWhiteCache(t *testing.T) {
	err := initNewRedis()
	if err != nil {
		return
	}

	//assert.Equal(t, 1, len(rds.GetTopKey()))

	v1Res, _ := rds.Get(context.Background(), "test_white_key")
	assert.Equal(t, "test_white_value", v1Res)

	//delCount, err := rds.Del(context.Background(), key1, key2)
	//assert.Nil(t, err)
	//assert.Equal(t, int64(2), delCount)
}
