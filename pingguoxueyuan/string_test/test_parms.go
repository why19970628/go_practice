package main

import "fmt"

func main() {
	key := "info_1"
	params := []interface{}{key}
	params = append(params, "777", "test")
	params = append(params, "888", "test2")
	fmt.Println(params)
}
