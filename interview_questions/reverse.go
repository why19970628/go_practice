package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func reverse(n int) {
	var (
		flag      int // 1正 -1负
		resultStr string
	)

	if n > 0 {
		flag = 1
	} else if n < 0 {
		flag = -1
	} else {
		log.Fatal("n == 0")
	}
	str := strconv.Itoa(n)
	for i := 0; i < len(str)-1; i++ {
		resultStr = resultStr + string(str[len(str)-1-i])
	}
	if n == 0 {
		log.Println("n == 0 ")
		return
	}
	resultInt, err := strconv.Atoi(resultStr)
	if err != nil || resultInt > math.MaxInt32 {
		fmt.Println(err)
	}
	fmt.Println(resultInt * flag)

}

func main() {

	n := -22121312310
	reverse(n)
}
