package main

import "fmt"

// ［2, 1,3,6, 5,7,8］ golang实现

func main() {
	input := [][]int{{2, 3, 6}, {1, 3, 5, 8}, {1, 6, 7}}
	output := mergeAndDeduplicate(input)
	fmt.Println(output)
}

func mergeAndDeduplicate(queues [][]int) []int {
	result := make([]int, 0)
	visited := make(map[int]bool)
	maxLen := 0
	for i := 0; i < len(queues); i++ {
		maxLen = max(maxLen, len(queues[i]))
	}
	for i := 0; i < maxLen; i++ {
		for _, list := range queues {
			if len(list)-1 >= i {
				val := list[i]
				if !visited[val] {
					result = append(result, val)
					visited[val] = true
				}
			}
		}
	}

	return result
}
