package main

import "container/list"

/**
102. 二叉树的层序遍历
*/
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val)
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []int{}          //清空层的数据

	}

	return res
}

/**
102. 二叉树的层序遍历：使用切片模拟队列，易理解
迭代
https://programmercarl.com/0102.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%B1%82%E5%BA%8F%E9%81%8D%E5%8E%86.html
*/
func levelOrderWithSlice(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	curLevel := []*TreeNode{root}

	for len(curLevel) > 0 {
		tempArr := []*TreeNode{}

		tempLevel := []int{}
		length := len(curLevel)

		for i := 0; i < length; i++ {
			node := curLevel[i]
			if node.Left != nil {
				tempArr = append(tempArr, node.Left)
			}
			if node.Right != nil {
				tempArr = append(tempArr, node.Right)
			}
			tempLevel = append(tempLevel, node.Val)
		}
		res = append(res, tempLevel)
		curLevel = tempArr
	}

	return
}
