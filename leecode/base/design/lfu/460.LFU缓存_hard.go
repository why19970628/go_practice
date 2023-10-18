package lfu

/*
https://leetcode.cn/problems/lfu-cache/

*/

//type LFUCache struct {
//	cache               map[int]*Node
//	freq                map[int]*DoubleList
//	ncap, size, minFreq int
//}
//
//func Constructor(capacity int) LFUCache {
//	return LFUCache{
//		cache: make(map[int]*Node),
//		freq:  make(map[int]*DoubleList),
//		ncap:  capacity,
//	}
//}
//
//func (this *LFUCache) Get(key int) int {
//	if node, ok := this.cache[key]; ok {
//		this.IncFreq(node)
//		return node.val
//	}
//	return -1
//}
//
//func (this *LFUCache) Put(key, value int) {
//	if this.ncap == 0 {
//		return
//	}
//	if node, ok := this.cache[key]; ok {
//		node.val = value
//		this.IncFreq(node)
//	} else {
//		if this.size >= this.ncap {
//			node := this.freq[this.minFreq].RemoveLast()
//			delete(this.cache, node.key)
//			this.size--
//		}
//		x := &Node{key: key, val: value, freq: 1}
//		this.cache[key] = x
//		if this.freq[1] == nil {
//			this.freq[1] = CreateDL()
//		}
//		this.freq[1].AddFirst(x)
//		this.minFreq = 1
//		this.size++
//	}
//}
//
//func (this *LFUCache) IncFreq(node *Node) {
//	_freq := node.freq
//	this.freq[_freq].Remove(node)
//	if this.minFreq == _freq && this.freq[_freq].IsEmpty() {
//		this.minFreq++
//		delete(this.freq, _freq)
//	}
//
//	node.freq++
//	if this.freq[node.freq] == nil {
//		this.freq[node.freq] = CreateDL()
//	}
//	this.freq[node.freq].AddFirst(node)
//}
//
//type DoubleList struct {
//	head, tail *Node
//}
//
//type Node struct {
//	prev, next     *Node
//	key, val, freq int
//}
//
//func CreateDL() *DoubleList {
//	head, tail := &Node{}, &Node{}
//	head.next, tail.prev = tail, head
//	return &DoubleList{
//		head: head,
//		tail: tail,
//	}
//}
//
//func (this *DoubleList) AddFirst(node *Node) {
//	node.next = this.head.next
//	node.prev = this.head
//
//	this.head.next.prev = node
//	this.head.next = node
//}
//
//func (this *DoubleList) Remove(node *Node) {
//	node.prev.next = node.next
//	node.next.prev = node.prev
//
//	node.next = nil
//	node.prev = nil
//}
//
//func (this *DoubleList) RemoveLast() *Node {
//	if this.IsEmpty() {
//		return nil
//	}
//
//	last := this.tail.prev
//	this.Remove(last)
//
//	return last
//}
//
//func (this *DoubleList) IsEmpty() bool {
//	return this.head.next == this.tail
//}

type node struct {
	key, val, freq int
	prev, next     *node
}

/**
双向列表可以实现O(1)删除操作
*/
type doubleList struct {
	head, tail *node
}

type LFUCache struct {
	cap         int
	keyToNode   map[int]*node
	freqToNodes map[int]*doubleList
	minFreq     int
}

func Constructor(cap int) LFUCache {
	return LFUCache{
		cap:         cap,
		keyToNode:   map[int]*node{},
		freqToNodes: map[int]*doubleList{},
		minFreq:     0,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.keyToNode[key]; ok {
		//增加评率
		this.increaseFreq(node)
		return node.val
	}
	return -1
}

func (this *LFUCache) Put(key, val int) {
	/**
	步骤：
	1.判断容量是否合法
	2.判断key是否存在，是则更新val，并增加频率

	3.key不存在，新增key，需要判断容量是否还有
	3.1 否：删除最小使用频率的key
	3.2 是: 正常插入

	4.新增操作minFreq=1
	*/
	if this.cap <= 0 {
		return
	}

	if node, ok := this.keyToNode[key]; ok {
		node.val = val
		//增加评率
		this.increaseFreq(node)
		return
	}

	if len(this.keyToNode) >= this.cap {
		//删除最小频率的key
		this.deleteMinFreq()
	}

	//新增key操作
	value := &node{
		key:  key,
		val:  val,
		freq: 1,
	}
	this.keyToNode[key] = value

	if this.freqToNodes[value.freq] == nil {
		this.freqToNodes[value.freq] = doubleListConstructor()
	}
	this.freqToNodes[value.freq].add(value)

	this.minFreq = 1
}

/**
增加node频率
*/
func (this *LFUCache) increaseFreq(node *node) {
	/**
	DL:doubleList,双向链表
	步骤：
	1.先从以前DL中移除node
	2.判断以前freq对应DL是否为空,是则minFreq++
	3.node.freq++并加入到新的DL中
	*/
	beforeFreq := node.freq
	oldDL := this.freqToNodes[beforeFreq]
	oldDL.remove(node)

	if oldDL.isEmpty() && beforeFreq == this.minFreq {
		this.minFreq++
	}

	node.freq++
	if this.freqToNodes[node.freq] == nil {
		this.freqToNodes[node.freq] = doubleListConstructor()
	}
	this.freqToNodes[node.freq].add(node)
}

/**
删除最小频率node
*/
func (this *LFUCache) deleteMinFreq() {
	/**
	步骤：
	1.DL中需要删除
	2.keyToNode需要删除
	*/
	dl := this.freqToNodes[this.minFreq]
	last := dl.last()

	dl.remove(last)
	delete(this.keyToNode, last.key)
}

func doubleListConstructor() *doubleList {
	head, tail := &node{}, &node{}
	head.next, tail.prev = tail, head

	return &doubleList{
		head: head,
		tail: tail,
	}
}

func (this *doubleList) add(node *node) {
	prev, next := this.head, this.head.next
	prev.next, node.prev = node, prev
	next.prev, node.next = node, next
}

func (this *doubleList) remove(node *node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
	node.prev, node.next = nil, nil
}

func (this *doubleList) first() *node {
	return this.head.next
}

func (this *doubleList) last() *node {
	return this.tail.prev
}

func (this *doubleList) isEmpty() bool {
	return this.head.next == this.tail
}

//作者：kong
//链接：https://leetcode.cn/problems/lfu-cache/solutions/940944/hash-mapduo-shuang-xiang-lie-biao-you-xi-pi03/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
