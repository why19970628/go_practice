package main

import (
	"fmt"
)

func main()  {
	string := "dcs232王"
	//int, err := strconv.Atoi(string)
	for i, i2 := range string {
		fmt.Println(i, i2)
	}
}
