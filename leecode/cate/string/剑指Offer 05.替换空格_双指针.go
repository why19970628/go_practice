package main

import "fmt"

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1： 输入：s = "We are happy."
输出："We%20are%20happy."
*/

// 遍历添加
func replaceSpaceV1(s string) string {
	bt := []byte(s)
	res := make([]byte, 0)
	for i := 0; i < len(bt); i++ {
		if string(bt[i]) == " " {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, bt[i])
		}
	}
	return string(res)
}

// 原地修改 双指针
func replaceSpaceV2(s string) string {
	oriBt := []byte(s)
	oldN := len(oriBt)
	var spaceCount int
	for i := 0; i < len(oriBt); i++ {
		if oriBt[i] == ' ' {
			spaceCount++
		}
	}
	if spaceCount == 0 {
		return s
	}

	// 数组扩容
	temp := make([]byte, 2*spaceCount)
	oriBt = append(oriBt, temp...)

	leftIndex := oldN - 1
	rightIndex := len(oriBt) - 1

	for leftIndex > 0 {
		if oriBt[leftIndex] != ' ' {
			oriBt[rightIndex] = oriBt[leftIndex]
			rightIndex--
		} else {
			oriBt[rightIndex] = '0'
			oriBt[rightIndex-1] = '2'
			oriBt[rightIndex-2] = '%'
			rightIndex -= 3
		}
		leftIndex--
	}

	return string(oriBt)
}

func main() {
	fmt.Println(replaceSpaceV2("We are happy."))
}
