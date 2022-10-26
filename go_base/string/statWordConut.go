package main

import (
	"fmt"
	"strings"
)

func stat() {
	s := "how do you do"
	s1 := strings.Split(s, " ")
	var map_str map[string]int = make(map[string]int, 128)
	for _, v := range s1 {
		result, ok := map_str[v]
		if !ok {
			map_str[v] = 1
		} else {
			map_str[v] = result + 1
		}
	}
	fmt.Println(s1)
	fmt.Println(map_str)
}
func main() {
	stat()
}
