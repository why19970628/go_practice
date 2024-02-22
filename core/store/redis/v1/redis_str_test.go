package redis

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisStr(t *testing.T) {
	key1 := "1"
	key2 := "2"
	var (
		v1 = "123"
		v2 = 456
	)
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

//
//func TestRedisStrInDecr(t *testing.T) {
//	key1, _ := rds.NewRedisKey("business", "str_test_key_indecr")
//	var (
//		value int64 = 2
//	)
//	val1, _ := rds.Incr(key1)
//	assert.Equal(t, int64(1), val1)
//
//	val2, _ := rds.IncrBy(key1, value)
//	assert.Equal(t, int64(3), val2)
//
//	val3, _ := rds.Decr(key1)
//	assert.Equal(t, int64(2), val3)
//
//	val4, _ := rds.DecrBy(key1, value)
//	assert.Equal(t, int64(0), val4)
//	_, _ = rds.Del(key1)
//
//}
