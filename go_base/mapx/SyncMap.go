package mapx

import (
	"fmt"
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

func main1() {
	m := &M{
		Map:  make(map[interface{}]interface{}),
		lock: new(sync.Mutex),
	}
	m.Set("2", "100")
	fmt.Println(m)
	fmt.Println(m.Get("2"))
	fmt.Println(m.Get("3"))

}
