package backtrack

/*
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
*/
func combinationSum3(n, k int) [][]int {
	dfscombinationSum3(n, k, 1, 0)
	return res
}

func dfscombinationSum3(targetSum, n int, startIndex, sum int) {
	if len(path) == n {
		if sum == targetSum {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
	}

	for i := startIndex; i <= 9; i++ {
		// 枝剪
		// (9-i+1) < k-len(path)  剩余坑位不够
		//  sum+i > n  已经比结果大了, 继续无意义
		if sum+i > targetSum || (9-i+1) < n-len(path) {
			break
		}

		path = append(path, i)
		dfscombinationSum3(targetSum, n, startIndex+1, sum+i)
		path = path[:len(path)-1] //回溯
	}
}
