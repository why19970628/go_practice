package lru

/*
https://leetcode.cn/problems/lru-cache/description/
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

// 双向链表、HASH表
type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

type LRUCache struct {
	capacity int
	hashmap  map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		hashmap:  make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hashmap[key]; ok {
		remove(node)
		insert(this.head, node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.hashmap[key]; ok {
		node.val = value
		remove(node)
		insert(this.head, node)
	} else {
		node := &Node{
			key: key,
			val: value,
		}

		insert(this.head, node)
		this.hashmap[key] = node
		if len(this.hashmap) > this.capacity {
			last := this.tail.pre
			remove(last)
			delete(this.hashmap, last.key)
		}
	}
}

func insert(head, node *Node) {
	// head <-> next

	// head <-> node <-> next
	next := head.next

	// node <-> next
	node.next = next
	next.pre = node

	//  head <-> node
	head.next = node
	node.pre = head

}

func remove(node *Node) {
	// pre <-> node <-> next

	// pre -> next
	node.pre.next = node.next
	// pre <-> next
	node.next.pre = node.pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
