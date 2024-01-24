package main

func levelOrderV2(root *TreeNode) (resp [][]int) {
	if root == nil {
		return
	}
	level := []*TreeNode{}
	level = append(level, root)

	for len(level) > 0 {
		n := len(level)
		var tmp []int
		var t []*TreeNode
		for i := 0; i < n; i++ {
			cur := level[i]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				t = append(t, cur.Left)
			}
			if cur.Right != nil {
				t = append(t, cur.Right)
			}
		}
		level = t
		resp = append(resp, tmp)
	}
	return
}

func levelOrderV3(root *TreeNode) (resp [][]int) {
	if root == nil {
		return
	}
	level := []*TreeNode{}
	level = append(level, root)

	for len(level) > 0 {
		n := len(level)
		tmp := []int{}
		for i := 0; i < n; i++ {
			cur := level[0]
			level = level[1:]

			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				level = append(level, cur.Left)
			}
			if cur.Right != nil {
				level = append(level, cur.Right)
			}
		}

		resp = append(resp, tmp)
	}
	return
}
