package main

import "fmt"

// 1.
func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	var res int
	if len(arr) <= 1 {
		panic("arr length err")
	}
	var leftIndex = 0
	var rightIndex = 1
	for leftIndex < len(arr) && rightIndex < len(arr) && leftIndex < rightIndex {
		var sub = arr[rightIndex] - arr[leftIndex]
		if sub > res {
			res = sub
		}
		if sub <= 0 {
			leftIndex = rightIndex
		}
		rightIndex++
	}
	fmt.Println(res)
}

// 2.
type Admin struct {
	Id     int
	Name   string
	RoleId int
}

type Role struct {
	Id       int
	RoleName string
	PriId    string
}

type Privilege struct {
	Id     int
	Api    string
	Method string
	Thing  int // 事件
}

//3.
