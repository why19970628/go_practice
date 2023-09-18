package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
https://leetcode.cn/problems/top-k-frequent-elements/

给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。



示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]


*/

func topKFrequent_innerSort(nums []int, k int) []int {
	mp := make(map[int]int, 0)
	var resp []int
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	for key, _ := range mp {
		resp = append(resp, key)
	}
	fmt.Println("mp:", mp, resp)

	// 	核心思想：排序
	//	可以不用包函数，自己实现快排
	sort.Slice(resp, func(i, j int) bool {
		return mp[resp[i]] > mp[resp[j]]
	})
	fmt.Println("resp:", resp)
	return resp[:k]
}

func topKFrequent_quickSort(nums []int, k int) []int {
	mp := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	freqList := make([]int, 0, len(mp))
	for _, v := range mp {
		freqList = append(freqList, v)
	}

	quickSort(freqList, 0, len(freqList)-1)
	resp := make([]int, 0, k)
	for key, val := range mp {
		if val >= freqList[len(freqList)-k] {
			resp = append(resp, key)
		}
	}
	return resp
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	pIndex := partition(nums, l, r)
	quickSort(nums, pIndex+1, r)
	quickSort(nums, l, pIndex-1)
}

func partition(nums []int, l, r int) int {
	p := nums[r]
	i, j := l, l
	for j < r {
		if nums[j] < p {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

/*
桶排序代码实现

有多少个数字（此处数字已变成唯一）就创建多少个桶（空间换时间）
将不同频率的元素放入不同的桶中
比如桶1放入出现了1次的元素，桶2放入出现了2次的元素，以此类推。。。

从最大桶（倒序）K次逐步取出高频元素


https://leetcode.cn/problems/top-k-frequent-elements/solutions/2415909/gojie-fa-xiang-xi-zhu-shi-tong-pai-xu-or-n997/

*/

func topKFrequent(nums []int, k int) []int {
	// 确定每个数的频率
	frequence := make(map[int]int)
	for _, v := range nums {
		frequence[v]++
	}

	// 创建桶
	// 一共多少个数字就创建多少个桶, +1是为了对付index 0
	buckts := make([][]int, len(nums)+1)
	// 将不同频率的元素放入不同桶
	// 比如 桶1存放出现1次的元素，桶2存放出现2次的元素...
	for num, freq := range frequence {
		buckts[freq] = append(buckts[freq], num)
	}
	fmt.Println("buckts:", buckts)

	res := make([]int, 0)
	// 从最大桶(倒序）开始逐步取出高频元素
	for i := len(buckts) - 1; i >= 0; i-- {
		if k == 0 {
			break
		}
		for _, v := range buckts[i] {
			res = append(res, v)
			k--
			if k == 0 {
				break
			}
		}
	}
	return res
}

// 堆排序
type Item struct {
	// 元素本身
	Num int
	// 元素出现次数
	Count int
}

// 对象实现Heap的五个接口
type Heap []*Item

// 小顶堆，让大的"沉下去"，以便每次把小的直接Pop掉
func (h Heap) Less(i, j int) bool {
	return h[i].Count < h[j].Count
}

// 简单的交换
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 简单的长度计算
func (h Heap) Len() int {
	return len(h)
}

// 每次直接把头部干掉(小的)
func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:h.Len()-1]
	return x
}

// 入队
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func topKFrequent_heapSort(nums []int, k int) []int {
	// 创建对象
	h := &Heap{}
	// 初始化堆
	heap.Init(h)

	// 统计每个元素出现次数
	frequence := make(map[int]int, 0)
	// key是元素本身，value是元素对应出现次数
	for _, num := range nums {
		frequence[num]++
	}

	// 把所有元素（唯一）入队
	for key, value := range frequence {
		heap.Push(h, &Item{Num: key, Count: value})
		// 所以每次元素超过K个了，就出队列，避免冗余计算
		//if h.Len() > k {
		//	// 因为是小顶堆，所以能确保出现频率高的k个不会被出队
		//	heap.Pop(h)
		//}
	}
	for _, item := range *h {
		fmt.Println(item)
	}
	// 统计结果
	res := make([]int, k)
	for i := 0; i < k; i++ {
		// 当前队列里面就保存的k个频率最高的元素，题目没有要求按照什么顺序排列，所以直接赋值
		res[i] = heap.Pop(h).(*Item).Num
		// 如果较真一些，按照出现频率从大到小来排列（即倒序），就可以写成
		// res[k-i-1] = heap.Pop(h).(*Item).Num
	}
	return res
}

func main() {
	fmt.Println(topKFrequent_heapSort([]int{1, 3, 4, 4, 3, 3, 2}, 2))
}
