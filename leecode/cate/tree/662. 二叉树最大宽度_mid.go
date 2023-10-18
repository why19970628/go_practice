package main

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	resp := 0
	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		tempArr := []*TreeNode{}

		tempLevel := []int{}
		length := len(curLevel)

		for i := 0; i < length; i++ {
			node := curLevel[i]
			if node == nil {
				continue
			}

			if node.Left != nil {
				tempArr = append(tempArr, node.Left)
			} else {
				tempArr = append(tempArr, nil)
			}
			if node.Right != nil {
				tempArr = append(tempArr, node.Right)
			}
			tempLevel = append(tempLevel, node.Val)
		}
		resp = max4(resp, len(tempLevel))
		curLevel = tempArr
	}

	return resp
}

func max4(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*

获取二叉树的宽度想当然的就会想到使用层序遍历，但注意到题目中要求的宽度是包含空节点的，那么就需要记录层中节点在该层中的位置。
具体做法是用一个队列记录每层的节点通过BFS向下遍历，这里用了一个map记录节点在层中的位置（也可以用结构体和节点绑定），记root节点位置为1。
可以发现左节点位置是父节点位置2-1，右节点位置是父节点位置2，最后将每层最右边的位置-最左边的位置+1即是该层的层宽


作者：MythicMyuu
链接：https://leetcode.cn/problems/maximum-width-of-binary-tree/solutions/1363608/golang-by-mythicmyuu-sjg5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func widthOfBinaryTree2(root *TreeNode) int {
	ans := 0
	mp := map[*TreeNode]int{}
	q := []*TreeNode{root}
	mp[root] = 1
	for len(q) > 0 {
		Len := len(q)
		ans = max(ans, mp[q[len(q)-1]]-mp[q[0]]+1)
		for i := 0; i < Len; i++ {
			node := q[0]
			q = q[1:]
			v := mp[node]
			if node.Left != nil {
				// 思路不难，无非就是给每个节点编上编号，
				//如果父节点的编号为 i,那么其左子节点的编号为 2i，右子节点的编号为2i+1。
				q = append(q, node.Left)
				mp[node.Left] = v * 2
			}
			if node.Right != nil {
				q = append(q, node.Right)
				mp[node.Right] = v*2 + 1
			}
		}
	}
	return ans
}

//
//作者：flames_c
//链接：https://leetcode.cn/problems/maximum-width-of-binary-tree/solutions/2345098/guan-fang-ti-jie-jiu-xiang-bian-bian-by-vza8p/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
