package main

//
//import "fmt"
//
//// 1.
//func main() {
//	arr := []int{7, 1, 5, 3, 6, 4}
//	var res int
//	if len(arr) <= 1 {
//		panic("arr length err")
//	}
//	var leftIndex = 0
//	var rightIndex = 1
//	for leftIndex < len(arr) && rightIndex < len(arr) && leftIndex < rightIndex {
//		var sub = arr[rightIndex] - arr[leftIndex]
//		if sub > res {
//			res = sub
//		}
//		if sub <= 0 {
//			leftIndex = rightIndex
//		}
//		rightIndex++
//	}
//	fmt.Println(res)
//}
//
//// 2.
//type Admin struct {
//	Id     int
//	Name   string
//	RoleId int
//}
//
//type Role struct {
//	Id       int
//	RoleName string
//	PriId    string
//}
//
//type Privilege struct {
//	Id     int
//	Api    string
//	Method string
//	Thing  int // 事件
//}
//
////3.
//
//type TreeNode struct {
//	Left  *TreeNode
//	Right *TreeNode
//	data  int
//}
//
//var arr2 []int
//var sum = 22
//var hasPath bool // result
//
//func (node *TreeNode) TraverseFunc(f func(node *TreeNode)) {
//	if node == nil {
//		return
//	}
//	arr2 = append(arr2, node.data)
//	a := 0
//	for _, v := range arr2 {
//		a += v
//		if a == sum {
//			hasPath = true
//		}
//	}
//
//	node.Left.TraverseFunc(f)
//	f(node)
//	node.Right.TraverseFunc(f)
//}
