package main

import (
	"fmt"
	"slices"
)

/*
https://leetcode.cn/problems/leaf-similar-trees/description/?envType=study-plan-v2&envId=leetcode-75

请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
*/
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	resp1 := rootList(root1)
	resp2 := rootList(root2)
	fmt.Println(resp1, resp2)
	return slices.Equal(resp1, resp2)
}

func rootList(root1 *TreeNode) []int {
	resp := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Right == nil && root.Left == nil {
			resp = append(resp, root.Val)
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root1)
	return resp
}
