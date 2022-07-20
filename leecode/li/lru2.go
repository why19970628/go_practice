package main

import "container/list"

// LRU是Least Recently Used的缩写，即最近最少使用，是一种常用的缓存淘汰算法，选择最近最久未使用的数据予以淘汰。
//该算法赋予每个数据一个访问字段，用来记录一个元素自上次被访问以来所经历的时间 t，当须淘汰一个数据时，选择现有数据中其 t 值最大的，即最近最少使用的页面予以淘汰。
//这里我们使用链表数据结构，可以省去设置访问时间t字段。
//我们可以将经常访问的数据插入链表头，而链表尾部的数据自然就成为最久未访问的数据。

type keyLru struct {
	limit    int                      //缓存数量
	evicts   *list.List               //双向链表用于淘汰数据
	elements map[string]*list.Element //记录缓存数据
	onEvict  func(key string)         //数据淘汰时回调函数，对数据做一些处理
}

func NewKeyLru(limit int, onEvict func(key string)) *keyLru {
	return &keyLru{
		limit:    limit,
		evicts:   list.New(),
		elements: make(map[string]*list.Element),
		onEvict:  onEvict,
	}
}

// 添加元素到缓存
func (klru *keyLru) add(key string) {
	//判断添加到值是否存在于缓存中
	if elem, ok := klru.elements[key]; ok {
		klru.evicts.MoveToFront(elem)
		return
	}
	// 添加新节点
	elem := klru.evicts.PushFront(key)
	klru.elements[key] = elem

	// 检查是否超出容量，如果超出就淘汰末尾节点数据
	if klru.evicts.Len() > klru.limit {
		klru.removeOldest()
	}
}

// 淘汰末尾节点
func (klru *keyLru) removeOldest() {
	elem := klru.evicts.Back() //获取链表末尾节点
	if elem != nil {
		klru.removeElement(elem)
	}
}

//删除节点操作
func (klru *keyLru) removeElement(e *list.Element) {
	// 删除末尾元素
	klru.evicts.Remove(e)

	key := e.Value.(string)
	// 删除缓存中数据
	delete(klru.elements, key)

	klru.onEvict(key)
}
