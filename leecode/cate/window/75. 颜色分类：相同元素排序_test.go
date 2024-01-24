package window

/*
https://leetcode.cn/problems/sort-colors/

给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。

*/

// map
func sortColors(nums []int) {
	var hash [3]int
	for _, color := range nums {
		hash[color]++
	}
	color := 0
	for i := 0; i < len(nums); i++ {
		for hash[color] == 0 {
			color++
		}
		hash[color]--
		nums[i] = color
	}
}

// 双指针
func sortColorsV3(nums []int) {
	n := len(nums)
	c0, c2 := 0, n-1
	for i := 0; i <= c2; {
		if nums[i] == 0 && i != c0 {
			nums[i], nums[c0] = nums[c0], nums[i]
			c0++
		} else if nums[i] == 2 && i != c2 {
			nums[i], nums[c2] = nums[c2], nums[i]
			c2--
		} else {
			i++
		}
	}
}

func sortColorsV2(nums []int) {
	c0, c2 := 0, len(nums)-1
	for i := 0; i <= c2; {
		if nums[i] == 0 && i != c0 {
			nums[i], nums[c0] = nums[c0], nums[i]
			c0++
		} else if nums[i] == 2 && i != c2 {
			nums[i], nums[c2] = nums[c2], nums[i]
			c2--
		} else {
			i++
		}
	}

}

func sortColorsV4(nums []int) {
	var hash [3]int
	for _, color := range nums {
		hash[color]++
	}

	index := 0
	for key, count := range hash {
		for count > 0 {
			nums[index] = key
			index++
			count--
		}
	}
}
