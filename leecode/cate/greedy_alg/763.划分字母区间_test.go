package greedy_alg

func partitionLabels(s string) []int {
	res := []int{}
	marks := [26]int{}
	for index, v := range s {
		marks[v-'a'] = index
	}
	right := 0
	left := 0
	for index, v := range s {
		// 计算该区间内，最大的边界
		right = max(right, marks[v-'a'])
		if index == right {
			res = append(res, right-left+1)
			left = index + 1
		}
	}
	return res
}
