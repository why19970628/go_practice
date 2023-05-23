package main

import "fmt"

func strStr(a, b string) bool {
	abt := []byte(a)
	bbt := []byte(b)

	an := len(abt)
	bn := len(bbt)

	for i := 0; i < an; i++ {
		fmt.Println("i", i)
		var is bool
		for j := 0; j < bn; j++ {
			fmt.Println("---", j, string(abt[i+j]), string(bbt[j]))
			if abt[i+j] != bbt[j] {
				break
			}
			if j == bn-1 {
				is = true
			}
		}
		if is {
			return true
		}

	}
	return false
}

func main() {
	fmt.Println(strStr("leetcode", "leet"))

}
