package redis

import "time"

// openTopK 热点检测
func (r *Redis) openTopK() bool {
	return r.topK != nil
}

// openLocalCache 本地缓存
func (r *Redis) openLocalCache() bool {
	return r.localCache != nil
}

func (r *Redis) getFromLocalCache(key string) (interface{}, bool) {
	if r.localCache != nil {
		resp, err := r.localCache.Get(key)
		if err != nil || resp == nil {
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

		if r.inWhileList(key) {
			_ = r.localCache.SetWithExpire(key, resp, expiration)
			return
		}
	}
}

// inWhileList 白名单
func (r *Redis) inWhileList(key string) bool {
	if r.whiteList != nil {
		if _, ok := r.whiteList[key]; ok {
			return true
		}
	}
	return false
}
