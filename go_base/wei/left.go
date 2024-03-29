package main

import "fmt"

// <<        左位    双目运算，把<<左侧的运算数的各二进位全部左移若干位，高位丢去，低位补0，左移n位就是乘以2的n次方
// >>        右移    双目运算，把>>左侧的运算数的各二进制全部右移若干位，右移n位就是除以2的n次方
func main() {
	//右边空出的位用0填补
	i := 4 << 1
	fmt.Println(i == 8)
	//步骤：
	//4:0000 0100
	//4<<1:0000 1000
	//0000 1000:8

	//左边空出的位用0或者1填补
	b := 4 >> 1
	fmt.Println(b == 2)
	//步骤
	//4:0000 0100
	//4>>1:0000 0010
	//0000 0010:2

}
