package main

import "fmt"

//定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func letterCombinations(digits string) (res []string) {
	if digits == "" {
		return
	}
	mp := make(map[byte]string)
	mp['2'] = "abc"
	mp['3'] = "def"
	mp['4'] = "ghi"
	mp['5'] = "jkl"
	mp['6'] = "mno"
	mp['7'] = "pqrs"
	mp['8'] = "tuv"
	mp['9'] = "wxyz"

	var recur func(digits string, count int, ans string)
	recur = func(digits string, count int, ans string) {
		fmt.Println(ans)
		if len(digits) == count {
			res = append(res, ans)
			return //递归终止条件
		}
		value := mp[digits[count]] // s = abc
		for i := 0; i < len(value); i++ {
			//递归
			recur(digits, count+1, ans+string(value[i]))
		}
	}

	recur(digits, 0, "")
	return
}

func main() {
	res := letterCombinations("2345")
	fmt.Println(res)
}
