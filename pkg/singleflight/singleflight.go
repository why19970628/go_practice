package singleflight

import (
	"sync"
)

type singleFlightX struct {
	lock *sync.Mutex
	flag map[string]*call
}

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

func (s *singleFlightX) Do(key string, fn func() (interface{}, error)) (val interface{}, err error) {
	s.lock.Lock()
	if c, ok := s.flag[key]; ok {
		// 如果有相同 key 的请求正在进行中，则等待它完成
		s.lock.Unlock()
		c.wg.Wait()

		s.lock.Lock()
		delete(s.flag, key)
		s.lock.Unlock()
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)

	s.flag[key] = c
	s.lock.Unlock()
	c.val, c.err = fn()
	c.wg.Done()

	// 删除时加锁
	s.lock.Lock()
	delete(s.flag, key)
	s.lock.Unlock()

	return c.val, c.err
}

func NewSingleFlightX() *singleFlightX {
	return &singleFlightX{
		lock: new(sync.Mutex),
		flag: make(map[string]*call),
	}
}
