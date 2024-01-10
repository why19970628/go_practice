package main

import "fmt"

/*
一个数初始是0，只通过+1和*2两种操作，最少几次操作能让这个数变成114514？
*/

/*
*2在二进制中就是后面加一个0（左移，位数多一位），类似于十进制*10 ；
0001 1011 1111 0101 0010在最左边的1后面有16位数，所以*2了16次；
+1的话就是1的个数，举个例子，从0开始：

	①+1变成 1；

	②左移变成 10；

	③+1变成 11；

	④左移变成110；

从上例可以看出，+1操作是穿插在*2（左移）的，所以有几个1就有几次+1 ；
二进制110=十进制6 ；
6=((0+1)*2+1)*2，操作顺序由括号内到括号外。
*/
func minOperations(target int) int {
	operations := 0

	for target > 0 {
		if target%2 == 0 {
			// 如果目标数是偶数，直接除以2
			target /= 2
		} else {
			// 如果目标数是奇数，执行减一操作
			target--
		}
		operations++
	}

	return operations
}

func minOperationsV3(target int) int {
	operations := 0

	for target > 0 {
		if target&1 == 1 {
			target--
		} else {
			target >>= 1
		}
		operations++
	}
	return operations
}
func minOperationsV2(target int) int {
	operations := 0

	for target > 0 {
		// 如果目标数为奇数，只能执行加一操作，因为无法通过乘二得到奇数
		if target&1 == 1 {
			target-- // 执行加一操作
		} else {
			target >>= 1 // 目标数为偶数，右移一位（相当于除以2）
		}
		operations++
	}

	return operations
}

func main() {
	fmt.Println("minOperations", minOperationsV2(114514))
}
