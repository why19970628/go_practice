package mapx

import (
	"sync"
)

type M struct {
	Map  map[interface{}]interface{}
	lock *sync.Mutex
}

func (m *M) Set(k interface{}, v interface{}) {
	m.lock.Lock()
	m.Map[k] = v
	m.lock.Unlock()
}

func (m *M) Get(k interface{}) interface{} {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.Map[k]
}
