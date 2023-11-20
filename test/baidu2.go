package main

import (
	"fmt"
	"strings"
)

/*
0<num<1亿
*/

var chineseDigits = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var arr = []string{"十", "百", "千", "万", "亿"}

func numToChinese(num int) string {
	if num == 0 {
		return chineseDigits[num]
	}
	var result string
	//count := 0
	for num > 0 {
		a := num % 10
		result = chineseDigits[a] + result
		//if count < len(arr) && num > 10 {
		//	result = arr[count] + result
		//	count++
		//}
		num = num / 10
	}
	result = strings.TrimSuffix(result, "零")
	result = strings.TrimPrefix(result, "零")
	return result
}

func NumsToString(num int) string {
	// 中文普通话读法
	mp := make(map[int]string)
	mp[0] = "零"
	mp[2] = "二"
	mp[3] = "三"
	if num < 10 {
		return mp[0]
	} else if num < 100 {
		a := num % 10
		b := num / 10
		return mp[b] + "十" + mp[a]
	} else if num < 1000 {
		a := num % 10
		b := num / 100
		c := num / 100
		fmt.Println(a, b, c)
		return mp[c] + "百" + mp[b] + "十" + mp[a]
	}
	return ""
}
func main() {
	fmt.Println(numToChinese(213300))
}
