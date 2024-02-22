package redis

import (
	"context"
	"errors"
)

func (r *Redis) SAdd(ctx context.Context, key string, members ...interface{}) (insertCount int64, err error) {
	if len(members) == 0 {
		err = errors.New("SAdd的参数不能少于1个")
		return
	}
	insertCount, err = r.Client.SAdd(ctx, key, members...).Result()
	return
}

func (r *Redis) SMembers(ctx context.Context, key string) (res []string, err error) {
	res, err = r.Client.SMembers(ctx, key).Result()
	return
}

func (r *Redis) SIsMember(ctx context.Context, key string, member interface{}) (isMember bool, err error) {
	isMember, err = r.Client.SIsMember(ctx, key, member).Result()
	return
}

func (r *Redis) SCard(ctx context.Context, key string) (memberCount int, err error) {
	val, err := r.Client.SCard(ctx, key).Result()
	memberCount = int(val)
	return
}

func (r *Redis) SDiff(ctx context.Context, keys ...string) (res []string, err error) {
	if len(keys) < 2 {
		res, err = []string{}, errors.New("SDiff的参数不能少于2")
		return
	}
	res, err = r.Client.SDiff(ctx, keys...).Result()
	return
}

func (r *Redis) SInter(ctx context.Context, keys ...string) (res []string, err error) {
	if len(keys) < 2 {
		res, err = []string{}, errors.New("SInter的参数不能少于2")
		return
	}
	res, err = r.Client.SInter(ctx, keys...).Result()
	return
}

func (r *Redis) SInterStore(ctx context.Context, destination string, keys ...string) (elemCount int64, err error) {
	if len(keys) < 2 {
		elemCount, err = 0, errors.New("SInterStore的参数不能少于2")
		return
	}
	elemCount, err = r.Client.SInterStore(ctx, destination, keys...).Result()
	return
}

func (r *Redis) SDiffStore(ctx context.Context, destination string, keys ...string) (res int64, err error) {
	if len(keys) < 2 {
		res, err = 0, errors.New("CustomClient的参数不能少于2")
		return
	}
	res, err = r.Client.SDiffStore(ctx, destination, keys...).Result()
	return
}

func (r *Redis) SMove(ctx context.Context, source, destination string, member interface{}) (ok bool, err error) {
	ok, err = r.Client.SMove(ctx, source, destination, member).Result()
	return
}

func (r *Redis) SPop(ctx context.Context, key string) (res string, err error) {
	res, err = r.Client.SPop(ctx, key).Result()
	return
}

func (r *Redis) SRandMember(ctx context.Context, key string) (res string, err error) {
	res, err = r.Client.SRandMember(ctx, key).Result()
	return
}

func (r *Redis) SRandMemberN(ctx context.Context, key string, count int64) (res []string, err error) {
	res, err = r.Client.SRandMemberN(ctx, key, count).Result()
	return
}

func (r *Redis) SRem(ctx context.Context, key string, members ...interface{}) (removeElemCount int64, err error) {
	if len(members) == 0 {
		err = errors.New("SRem的参数不能少于1个")
		return
	}
	removeElemCount, err = r.Client.SRem(ctx, key, members...).Result()
	return
}

func (r *Redis) SUnion(ctx context.Context, keys ...string) (res []string, err error) {
	if len(keys) < 2 {
		err = errors.New("SUnion的参数不能少于两个")
		return
	}
	res, err = r.Client.SUnion(ctx, keys...).Result()
	return
}

func (r *Redis) SUnionStore(ctx context.Context, destination string, keys ...string) (elemCount int64, err error) {
	if len(keys) < 2 {
		err = errors.New("SUnionStore的参数不能少于两个")
		return
	}
	elemCount, err = r.Client.SUnionStore(ctx, destination, keys...).Result()
	return
}

func (r *Redis) SSCan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	res, cursor, err := r.Client.SScan(ctx, key, cursor, match, count).Result()
	return res, cursor, err
}
