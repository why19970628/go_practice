package main

import (
	"fmt"
	"strconv"
)

func Process(num int)  {
	strNum := strconv.Itoa(num)
	for i := 0; i < len(strNum)/2; i++ {
		j := len(strNum ) -1
		if strNum[i] != strNum[j]{
			fmt.Printf("不是回文数")
			return
		}
		j--
	}
}

func main() {
	var num int = 54321
	Process(num)
}
