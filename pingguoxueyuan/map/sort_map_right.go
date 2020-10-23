package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	mapStr := make(map[string]int, 100)
	//mapStr["stu2"] = 2
	//mapStr["stu1"] = 1
	//mapStr["stu3"] = 3
	//mapStr["stu4"] = 4

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu_%02d", i)
		value := rand.Intn(100)
		mapStr[key] = value
	}


	fmt.Println(mapStr)

	v, ok := mapStr["hhh"]
	if ok{
		fmt.Println(v)
	}else {
		fmt.Println("no err")
	}

	array := make([]string, 0, 100)
	//fmt.Println(array)
	for k, _ := range mapStr {
		array = append(array, k)
	}
	sort.Strings(array)
	for _, i3 := range array {
		println(i3, mapStr[i3])
	}




}
