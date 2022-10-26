package main

import (
	"fmt"
	"strings"
)

func main() {
	names := [3]string{"Alex", "超级无敌小", "傻儿子"}
	fmt.Println(names)

	v3 := names[2]
	fmt.Println(v3)

	v1 := names[0]
	fmt.Println(strings.ToLower(v1))

	names[2] = "大烧饼"
	fmt.Println(names)

	ress := strings.Fields("a b c f d g")
	for k, v := range ress {
		fmt.Println(k, v)
	}

}
