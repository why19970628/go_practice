package main

import "container/heap"

/*


 */

/*
【Go】维护两个堆，一个大根堆保存数据流前半部分，一个小根堆保存数据流后半部分
https://leetcode.cn/problems/find-median-from-data-stream/solutions/2283271/go-wei-hu-liang-ge-dui-yi-ge-da-gen-dui-nq2hh/?envType=study-plan-v2&envId=top-100-liked
*/

type MinHeap295 []int

func (hp MinHeap295) Len() int           { return len(hp) }
func (hp MinHeap295) Less(i, j int) bool { return hp[i] < hp[j] }
func (hp MinHeap295) Swap(i, j int)      { hp[i], hp[j] = hp[j], hp[i] }

func (hp *MinHeap295) Push(x interface{}) {
	*hp = append(*hp, x.(int))
}

func (hp *MinHeap295) Pop() interface{} {
	n := len(*hp)
	x := (*hp)[n-1]
	*hp = (*hp)[:n-1]
	return x
}

type MaxHeap []int

func (hp MaxHeap) Len() int           { return len(hp) }
func (hp MaxHeap) Less(i, j int) bool { return hp[i] > hp[j] }
func (hp MaxHeap) Swap(i, j int)      { hp[i], hp[j] = hp[j], hp[i] }

func (hp *MaxHeap) Push(x interface{}) {
	*hp = append(*hp, x.(int))
}

func (hp *MaxHeap) Pop() interface{} {
	n := len(*hp)
	x := (*hp)[n-1]
	*hp = (*hp)[:n-1]
	return x
}

type MedianFinder struct {
	//维护后半个,数据小的放堆顶
	minhp *MinHeap295
	//维护前半个,数据大的放堆顶
	maxhp *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{&MinHeap295{}, &MaxHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	//如果是空的
	if this.minhp.Len() == 0 {
		heap.Push(this.minhp, num)
		return
	}
	//如果不为空
	//后半段的更多
	if this.minhp.Len() > this.maxhp.Len() {
		heap.Push(this.minhp, num)
		heap.Push(this.maxhp, heap.Pop(this.minhp).(int))
	} else {
		//一样多
		heap.Push(this.maxhp, num)
		heap.Push(this.minhp, heap.Pop(this.maxhp).(int))
	}
	// fmt.Println(this.maxhp)
	// fmt.Println(this.minhp)
	// fmt.Println()
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minhp.Len() > this.maxhp.Len() {
		return float64((*this.minhp)[0])
	} else {
		return float64((*this.minhp)[0]+(*this.maxhp)[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
