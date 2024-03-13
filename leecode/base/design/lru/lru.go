package lru

import (
	"container/list"
)

type data struct {
	key int
	val int
}
type lru struct {
	cap  int
	list *list.List
	mp   map[int]*list.Element
}

func (l *lru) Get(key int) int {
	val, ok := l.mp[key]
	if ok {
		l.list.MoveToFront(val)
		return val.Value.(*data).val
	}
	return -1
}

func (l *lru) Set(key, value int) {
	val, ok := l.mp[key]
	if ok {
		val.Value.(*data).val = value
		l.list.MoveToFront(val)
		return
	}
	newData := &data{
		key: key,
		val: value,
	}
	if l.list.Len() >= l.cap {
		delete(l.mp, l.list.Back().Value.(data).key)
		l.list.Remove(l.list.Back())
	}
	newElem := l.list.PushFront(newData)
	l.mp[key] = newElem

}
