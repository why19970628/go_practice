package main

func allPathsSourceTarget(graph [][]int) [][]int {
	ans := make([][]int, 0)
	n := len(graph)
	var dfs func([]int, int)
	dfs = func(ints []int, index int) {
		if index == n-1 {
			temp := make([]int, len(ints))
			copy(temp, ints)
			ans = append(ans, temp)
			return
		}
		for _, v := range graph[index] {
			ints = append(ints, v)
			dfs(ints, v)
			ints = ints[:len(ints)-1]
		}
	}
	dfs([]int{0}, 0)
	return ans
}

func allPathsSourceTargetV2(graph [][]int) [][]int {
	ans := make([][]int, 0)
	n := len(graph)
	var dfs func([]int, int)
	dfs = func(ints []int, index int) {
		if index == n-1 {
			temp := make([]int, len(ints))
			copy(temp, ints)
			ans = append(ans, temp)
			return
		}
		for _, v := range graph[index] {
			ints = append(ints, v)
			dfs(ints, v)
			ints = ints[:len(ints)-1]
		}
	}
	dfs([]int{0}, 0)
	return ans
}

func main() {
	graph := [][]int{
		{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	allPathsSourceTargetV2(graph)
}

//var (
//	allPathsSourceTargetResp [][]int
//	allPathsSourceTargetTemp []int
//)
//
//
//func allPathsSourceTargetDfs(n start int)  {
//	if
//}
