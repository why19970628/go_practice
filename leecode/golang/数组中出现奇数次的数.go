package main

import "fmt"

// 遍历数组，把数组中的元素作为一个map的key，判断这个key在map中是否已存在，若不存在value=1，
// 若已存在翻转value的值（若value=0翻转成1，若value=1翻转成0），最后value=1的key就是出现奇数次的数。
//
//func oddTimesNum(arr []int) {
//	timesMap := make(map[int]int)
//	for _, value := range arr {
//
//		v, ok := timesMap[value]
//		if ok {
//			if v == 0 {
//				timesMap[value] = 1
//			} else {
//				timesMap[value] = 0
//			}
//		} else {
//			timesMap[value] = 1
//		}
//	}
//	fmt.Println(timesMap)
//	for k, v := range timesMap {
//		if v == 1 {
//			fmt.Println(k)
//		}
//	}
//
//}

func a(arr []int) {
	m := make(map[int]int)
	for _, v := range arr {
		_, ok := m[v]
		if ok {
			m[v] += 1
			if m[v] == 2 {
				fmt.Println("m:", v, m[v])
			}
		} else {
			m[v] = 1
		}
	}
}
func main() {
	var arr []int
	arr = []int{1, 3, 2, 6, 5, 7, 8, 7, 2, 1}
	a(arr)
}
