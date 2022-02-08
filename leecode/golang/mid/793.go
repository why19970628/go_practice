package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, 0, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		var (
			a          bool
			highCount  int
			highCount2 int
		)

		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				highCount = highCount2
				a = true
				break
			}

			if temperatures[j] <= temperatures[i] {
				highCount2++
			}
		}

		fmt.Println("count", "a :")
		if a {
			if highCount == 0 {
				highCount++
			}
			res = append(res, highCount)
		} else {
			res = append(res, highCount)
		}

		fmt.Println("---")

	}

	return res

}

func dailyTemperatures2(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(arr))
}
