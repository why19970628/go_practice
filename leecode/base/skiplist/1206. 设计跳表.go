package main

type SkiplistNode struct {
	val  int
	node []*SkiplistNode
}

type Skiplist struct {
	Level int
	list  *SkiplistNode
}

func Constructor() Skiplist {
	return Skiplist{
		Level: -1,
		list: &SkiplistNode{
			val:  -1,
			node: make([]*SkiplistNode, 0),
		},
	}
}

func (this *Skiplist) Search(target int) bool {
	cur := this.list
	if cur.node == nil || len(cur.node) == 0 {
		return false
	}
	for i := this.Level; i >= 0; i-- {
		for cur.node[i] != nil && cur.node[i].val < target {
			cur = cur.node[i]
		}
	}
	cur = cur.node[0]
	return cur != nil && cur.val == target
}

func (this *Skiplist) Add(num int) {

}

func (this *Skiplist) Erase(num int) bool {
	return false
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
