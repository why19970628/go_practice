package main

import "fmt"

func sortMapValue(mp map[string]int) {
	reverseMap := make(map[int][]string)

	for key, value := range mp {
		fmt.Println(key, value)
		reverseMap[value] = append(reverseMap[value], key)
	}
	fmt.Println("reverseMap:", reverseMap)
}

func main() {
	map1 := make(map[string]int)
	map1["Mon"]=1
	map1["Tue"]=5
	map1["Wed"]=3
	map1["Thu"]=4
	map1["Fri"]=2
	map1["Sat"]=6
	map1["Sun"]=7
	fmt.Println(map1)

	sortMapValue(map1)
}

