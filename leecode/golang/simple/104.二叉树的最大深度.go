package main

func maxDepth(root *TreeNode) int {
	i := 0
	j := 0
	if root == nil {
		return 0
	}
	i = maxDepth(root.Left)
	j = maxDepth(root.Right)
	if i < j {
		i, j = j, i
	}
	return i + 1
}

func maxDepth2(root *TreeNode) int {
	var ans int = 0
	var f func(root *TreeNode, cnt int)
	f = func(root *TreeNode, cnt int) {
		if root == nil {
			return
		}
		cnt += 1
		if cnt > ans {
			ans = cnt
		}
		f(root.Left, cnt)
		f(root.Right, cnt)
	}
	f(root, 0)
	return ans
}

func main() {

}
