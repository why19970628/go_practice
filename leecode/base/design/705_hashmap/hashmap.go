package _05_hashmap

/*
不使用任何内建的哈希表库设计一个哈希集合（HashSet）。

实现 MyHashSet 类：

void add(key) 向哈希集合中插入值 key 。
bool contains(key) 返回哈希集合中是否存在这个值 key 。
void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。

0 <= key <= 106
最多调用 104 次 add、remove 和 contains
*/

//type MyHashSet struct {
//	arr []int
//}
//
//func Constructor() MyHashSet {
//	return MyHashSet{
//		arr: make([]int, 1000001),
//	}
//}
//
//func (this *MyHashSet) Add(key int) {
//	this.arr[key] = 1
//}
//
//func (this *MyHashSet) Remove(key int) {
//	this.arr[key] = 0
//}
//
//func (this *MyHashSet) Contains(key int) bool {
//	return this.arr[key] == 1
//
//}

// https://leetcode.cn/problems/design-hashset/solutions/343836/shuang-bai-can-kao-redis-wei-cao-zuo-yuan-li-yong-/

type MyHashSet struct {
	bitSet []uint64
}

func Constructor() MyHashSet {
	return MyHashSet{
		bitSet: make([]uint64, 0),
	}
}

func (this *MyHashSet) Add(key int) {
	bit := key % 64
	length := key / 64
	for length >= len(this.bitSet) {
		this.bitSet = append(this.bitSet, 0)
	}
	this.bitSet[length] = this.bitSet[length] | (1 << bit)
}

func (this *MyHashSet) Remove(key int) {
	bit := key % 64
	length := key / 64
	for length >= len(this.bitSet) {
		return
	}
	this.bitSet[length] = this.bitSet[length] & ^(1 << bit)
}

func (this *MyHashSet) Contains(key int) bool {
	bit := key % 64
	length := key / 64
	if length >= len(this.bitSet) {
		return false
	}
	return this.bitSet[length]&(1<<bit) != 0
}
