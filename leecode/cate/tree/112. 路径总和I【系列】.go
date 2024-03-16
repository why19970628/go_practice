package main

// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

// https://leetcode.cn/problems/path-sum/

// 递归
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val                                        // 将targetSum在遍历每层的时候都减去本层节点的值
	if root.Left == nil && root.Right == nil && targetSum == 0 { // 如果剩余的targetSum为0, 则正好就是符合的结果
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum) // 否则递归找
}

// 层序遍历模板
func hasPathSumV2(root *TreeNode, targetSum int) bool {
	//判断特殊情况根节点为空返回false
	if root == nil {
		return false
	}
	quene := []*TreeNode{root}
	for len(quene) > 0 {
		node := quene[0]
		quene = quene[1:]
		if node.Left == nil && node.Right == nil && node.Val == targetSum {
			return true
		}
		if node.Left != nil {
			node.Left.Val += node.Val
			quene = append(quene, node.Left)
		}
		if node.Right != nil {
			node.Right.Val += node.Val
			quene = append(quene, node.Right)
		}

	}
	return false
}

func hasPathSumDFS(root *TreeNode, targetSum int) bool {
	targetSum -= root.Val // 将targetSum在遍历每层的时候都减去本层节点的值
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSumDFS(root.Left, targetSum) || hasPathSumDFS(root.Right, targetSum)
}

//
//作者：可以是小盖
//链接：https://leetcode.cn/problems/path-sum/solutions/2054682/ke-yi-shi-xiao-gai-goyu-yan-zhu-xing-zhu-e40d/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
