package main

import "fmt"

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func levelOrder(root *TreeNode2) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var queue = []*TreeNode2{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		fmt.Println(counter)
		res = append(res, []int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			fmt.Println(res, level)
			queue = queue[1:]
		}
		level++
	}
	return res
}

//
//
//type data struct {
//	node  *TreeNode
//	floor int
//}
//
//
//func levelOrder2(root *TreeNode) [][]int {
//	// write code here
//	ans := [][]int{}
//	if root == nil {
//		return ans
//	}
//	stack := []data{}
//	stack = append(stack, data{root, 0})
//
//	tmp := []int{}
//	for len(stack) != 0 {
//		//取出头部
//		nowNode := stack[0]
//		stack = stack[1:]
//		//尾部添加
//		if nowNode.node.Left != nil {
//			stack = append(stack, data{nowNode.node.Left, nowNode.floor + 1})
//		}
//		if nowNode.node.Right != nil {
//			stack = append(stack, data{nowNode.node.Right, nowNode.floor + 1})
//		}
//		//将当前节点Val添入tmp
//		tmp = append(tmp, nowNode.node.Val)
//		if len(stack) == 0 || stack[0].floor != nowNode.floor {
//			//因为使用的是同一个数组，地址相同，直接添加可能会后面的更新前面的值
//			//因此需要每次声明一个局部变量，添加进去
//			tmpData:=[]int{}
//			tmpData=append(tmpData, tmp...)
//			ans = append(ans, tmpData)
//			//tmp添加进答案后重置
//			tmp = tmp[:0]
//		}
//	}
//	return ans
//}
//

var res [][]int

//func levelOrder3(root *TreeNode2) [][]int {
//	res = [][]int{}
//	dfs(root, 0)
//	return res
//}
//
//func dfs(root *TreeNode2, level int) {
//	fmt.Println("level:", level, res)
//	if root != nil {
//		if len(res) == level {
//			res = append(res, []int{})
//		}
//		res[level] = append(res[level], root.Val)
//		dfs(root.Left, level+1)
//		dfs(root.Right, level+1)
//	}
//}

func levelOrder4(root *TreeNode2) [][]int {
	res = [][]int{}
	solution(root, 0)
	return res

}

func solution(root *TreeNode2, level int) {
	if root != nil {
		// 第一次
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		solution(root.Left, level+1)
		solution(root.Right, level+1)
	}
}

func main() {
	node1 := &TreeNode2{3, nil, nil}
	node1.Left = &TreeNode2{9, nil, nil}
	node1.Right = &TreeNode2{20, nil, nil}
	node1.Right.Left = &TreeNode2{15, nil, nil}
	node1.Right.Right = &TreeNode2{7, nil, nil}
	fmt.Println(res)

	res2 := levelOrder4(node1)
	fmt.Printf("res is: %v\n", res2)
}
