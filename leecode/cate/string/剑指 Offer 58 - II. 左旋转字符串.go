package main

import "fmt"

// 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

// 输入: s = "abcdefg", k = 2
//输出: "cdefgab"

/*

为了让本题更有意义，提升一下本题难度：不能申请额外空间，只能在本串上操作。
不能使用额外空间的话，模拟在本串操作要实现左旋转字符串的功能还是有点困难的。
具体步骤为：

反转区间为前n的子串
反转区间为n到末尾的子串
反转整个字符串
*/
func reverseLeftWords(s string, n int) string {
	bt := []byte(s)
	length := len(bt) - 1
	reverse222(bt, 0, n-1)
	// n-1已经排过了，所以这里要从n开始
	reverse222(bt, n, length)
	reverse222(bt, 0, length)
	return string(bt)
}

// 切片是引用传递
func reverse222(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func main() {
	// 输入: s = "lrlose umgh", k = 6
	//输出: "umghlrlose"
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
}
