package main

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	Head, Tail *DLinkNode
}

type DLinkNode struct {
	key, value int
	Pre, Next  *DLinkNode
}

func InitDlinkNode(key, value int) *DLinkNode {
	return &DLinkNode{key, value, nil, nil}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		0,
		capacity,
		map[int]*DLinkNode{},
		InitDlinkNode(0, 0),
		InitDlinkNode(0, 0),
	}
	l.Head.Next = l.Tail
	l.Tail.Pre = l.Head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.UpdateToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := InitDlinkNode(key, value)
		for this.size >= this.capacity {
			this.DeleteLast()
		}
		this.cache[key] = node
		this.InsertNewHead(node)
	} else {
		node := this.cache[key]
		node.value = value
		this.UpdateToHead(node)
	}
}

func (this *LRUCache) UpdateToHead(node *DLinkNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	temp := this.Head.Next
	this.Head.Next = node
	node.Pre = this.Head
	node.Next = temp
	temp.Pre = node

}

func (this *LRUCache) DeleteLast() {
	node := this.Tail.Pre
	this.Tail.Pre = node.Pre
	node.Pre.Next = node.Next
	node.Pre = nil
	node.Next = nil
	this.size--
	delete(this.cache, node.key)
}

func (this *LRUCache) InsertNewHead(node *DLinkNode) {
	temp := this.Head.Next
	this.Head.Next = node
	node.Pre = this.Head
	temp.Pre = node
	node.Next = temp
	this.size++
}
