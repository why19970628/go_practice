package backtrack

/*

https://leetcode.cn/problems/word-search/description/

给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

*/

// https://leetcode.cn/problems/word-search/solutions/2031439/by-wwsw-jn43/
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	var canFind func(r, c, i int) bool
	canFind = func(r, c, i int) bool {
		if i == len(word) {
			return true
		} // word[:i]的每个字符均匹配，则找到了答案
		if r < 0 || r >= m || c < 0 || c >= n {
			return false
		} // 越界
		if used[r][c] {
			return false
		} // 题意说不允许重复访问某网格
		if board[r][c] != word[i] {
			return false
		} // 此路不通

		used[r][c] = true // 标记本网格已走过，防止本网格被重复走
		matchRest := canFind(r+1, c, i+1) ||
			canFind(r-1, c, i+1) ||
			canFind(r, c+1, i+1) ||
			canFind(r, c-1, i+1) // 四个方向探索
		if matchRest {
			return true
		} // 四条路之一找到了
		used[r][c] = false // 回溯：因本网格不能构建出目标路径，需要撤销此选择而去尝试别的选择

		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && canFind(i, j, 0) { // 根据word[0]找到起点，开始探索
				return true
			}
		}
	}
	return false
}
