package sync

import (
	"sync"
)

type singleton struct {
}

var once sync.Once

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		once.Do(func() {
			instance = &singleton{}
		})
	}
	return instance
}

//func (o *Once) Do(f func()) {
// Note: Here is an incorrect implementation of Do:
// 如果同时进行两次呼叫，则cas的获胜者将调用f(), 失败者会立即返回！！！,没有等待第一个对f的调用完成。
//	if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
//		f()
//	}
//
// Do guarantees that when it returns, f has finished.
// This implementation would not implement that guarantee:
// given two simultaneous calls, the winner of the cas would
// call f, and the second would return immediately, without
// waiting for the first's call to f to complete.
// This is why the slow path falls back to a mutex, and why
// the atomic.StoreUint32 must be delayed until after f returns.

// 先判断是否被操作过
// 会先判断 done 是否为 0，如果不为 0 说明还没执行过，就进入 doSlow
//if atomic.LoadUint32(&o.done) == 0 {
// Outlined slow-path to allow inlining of the fast-path.
//o.doSlow(f)
//}
//}

//func (o *Once) doSlow(f func()) {
//  互斥锁保证进程内顺序,
// 当中使用了互斥锁来保证只会执行一次,且顺序执行
//	o.m.Lock()
//	defer o.m.Unlock()
//	if o.done == 0 {
//		defer atomic.StoreUint32(&o.done, 1)
//		f()
//	}
//}
