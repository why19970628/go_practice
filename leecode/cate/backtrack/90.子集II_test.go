package backtrack

import (
	"fmt"
	"sort"
	"testing"
)

func subsetsWithDup(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)
	//res = append(res, []int{})
	sort.Ints(nums)
	subsetsWithDupDfs(nums, 0)
	return res
}

func subsetsWithDupDfs(nums []int, start int) {
	temp := make([]int, len(path))
	copy(temp, path)
	res = append(res, temp)

	for i := start; i < len(nums); i++ {
		// 同一树层上重复取2 就要过滤掉，同一树枝上就可以重复取2，因为同一树枝上元素的集合才是唯一子集！
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		subsetsWithDupDfs(nums, i+1)
		path = path[:len(path)-1]
	}
}

func TestSubsetsWithDup(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
