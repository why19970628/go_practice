package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	resp := make([][]int, 0)
	list := []*TreeNode{root}
	for len(list) > 0 {
		temp := make([]*TreeNode, 0)
		arr := make([]int, 0)
		for i := 0; i < len(list); i++ {
			cur := list[i]
			arr = append(arr, cur.Val)
			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}
		}
		list = temp
		if len(arr) > 0 {
			resp = append(resp, arr)
		}
	}
	l := 0
	r := len(resp) - 1
	for l < r {
		resp[l], resp[r] = resp[r], resp[l]
		l++
		r--
	}
	return resp
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(levelOrderBottom(root))
}
