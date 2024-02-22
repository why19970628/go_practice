package redis

import (
	"github.com/go-kratos/aegis/topk"
	"time"
)

// openTopK 热点检测
func (r *Redis) openTopK() bool {
	return r.topK != nil
}

// openLocalCache 本地缓存
func (r *Redis) openLocalCache() bool {
	return r.localCache != nil
}

func (r *Redis) getFromLocalCache(key string) (interface{}, bool) {
	if r.openLocalCache() {
		resp, _ := r.localCache.Get(key)
		if resp == nil {
			if val, ok := r.inWhileList(key); ok {
				return val, ok
			}
			return nil, false
		}
		return resp, true
	}
	return nil, false
}

func (r *Redis) updateLocalCache(key string, resp interface{}, expiration time.Duration) {
	if r.openLocalCache() {
		if r.openTopK() {
			if _, ok := r.topK.Add(key, 1); ok {
				_ = r.localCache.SetWithExpire(key, resp, expiration)
				return
			}
		}

		if _, ok := r.inWhileList(key); ok {
			_ = r.localCache.SetWithExpire(key, resp, expiration)
			return
		}
	}
}

// inWhileList 白名单
func (r *Redis) inWhileList(key string) (interface{}, bool) {
	if r.whiteList != nil {
		if val, ok := r.whiteList[key]; ok {
			return val, true
		}
	}
	return nil, false
}

func (r *Redis) GetTopKey() []topk.Item {
	return r.topK.List()
}
