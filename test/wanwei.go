package main

import (
	"fmt"
	"sync"
	"time"
)

// singleflight
// 缓存过期
// 设计结构体

type Cache struct {
	lock *sync.Mutex
	mp   map[string]CacheInfo
}

type CacheInfo struct {
	Expire int64
	Data   interface{}
}

func NewCache() *Cache {
	return &Cache{
		lock: new(sync.Mutex),
		mp:   make(map[string]CacheInfo),
	}
}

func (c *Cache) Set(key string, val interface{}) {
	c.lock.Lock()
	c.mp[key] = CacheInfo{
		Data: val,
	}
	c.lock.Unlock()
	return
}

func (c *Cache) SetWithExpire(key string, val interface{}, Expire int64) {
	c.lock.Lock()
	c.mp[key] = CacheInfo{
		Expire: time.Now().Unix() + Expire,
		Data:   val,
	}
	c.lock.Unlock()
	return
}

func (c *Cache) Get(key string) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	val, ok := c.mp[key]
	if ok {
		if val.Expire == 0 || val.Expire >= time.Now().Unix() {
			return val
		}
	}
	delete(c.mp, key)
	return nil
}

func main() {
	cache := NewCache()
	cache.Set("123", "VAL")
	fmt.Println(cache.Get("123"))

	cache.SetWithExpire("456", "VAL", 2)
	fmt.Println(cache.mp, time.Now().Unix())
	time.Sleep(time.Second * 3)
	fmt.Println(cache.mp, time.Now().Unix())
	fmt.Println(cache.Get("456"))

}
