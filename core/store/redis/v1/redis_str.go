package redis

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	err = r.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	r.updateLocalCache(key, value, expiration)
	return
}

func (r *Redis) SetWithNoExpiration(ctx context.Context, key string, value interface{}) (err error) {
	err = r.Client.Set(ctx, key, value, NoExpiration).Err()
	return
}

func (r *Redis) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (ok bool, err error) {
	ok, err = r.Client.SetNX(ctx, key, value, expiration).Result()
	return
}

func (r *Redis) Get(ctx context.Context, key string) (res interface{}, err error) {
	if r.openLocalCache() {
		resp, ok := r.getFromLocalCache(key)
		if ok {
			fmt.Println("hit local cache", key)
			return resp, nil
		}

	}
	res, err = r.Client.Get(ctx, key).Result()
	r.updateLocalCache(key, res, r.localCacheTTL)
	return
}

func (r *Redis) MSet(ctx context.Context, values ...interface{}) (err error) {
	if len(values) > 1000 {
		return errors.New("mset values数量不能超过1000")
	}
	err = r.Client.MSet(ctx, values).Err()
	return
}

func (r *Redis) MSetNX(ctx context.Context, values ...interface{}) (err error) {
	err = r.Client.MSetNX(ctx, values).Err()
	return
}

func (r *Redis) MGet(ctx context.Context, keys ...string) (res []interface{}, err error) {
	if len(keys) == 0 {
		err = errors.New("MGet的参数个数不能为0")
		return []interface{}{}, err
	}
	if len(keys) > 1000 {
		return []interface{}{}, errors.New("mget keys数量不能超过1000")
	}
	res, err = r.Client.MGet(ctx, keys...).Result()
	return
}

func (r *Redis) Incr(ctx context.Context, key string) (int64, error) {
	return r.Client.Incr(ctx, key).Result()
}

func (r *Redis) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return r.Client.IncrBy(ctx, key, value).Result()
}

func (r *Redis) Decr(ctx context.Context, key string) (int64, error) {
	return r.Client.Decr(ctx, key).Result()
}

func (r *Redis) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	return r.Client.DecrBy(ctx, key, value).Result()
}

func (r *Redis) GetRange(ctx context.Context, key string, start int64, end int64) (string, error) {
	return r.Client.GetRange(ctx, key, start, end).Result()
}

func (r *Redis) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	return r.Client.SetRange(ctx, key, offset, value).Result()
}

func (r *Redis) StrLen(ctx context.Context, key string) (int64, error) {
	return r.Client.StrLen(ctx, key).Result()
}

func (r *Redis) Append(ctx context.Context, key string, value string) (val int64, err error) {
	val, err = r.Client.Append(ctx, key, value).Result()
	return val, err
}

func (r *Redis) GetSet(ctx context.Context, key string, value string) (val string, err error) {
	val, err = r.Client.GetSet(ctx, key, value).Result()
	return
}
