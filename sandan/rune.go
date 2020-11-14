package sandan

import (
	"fmt"
	"strings"
)

func main() {
	urlSring := "dynamic/images/2020/03/28/A0950D2037B4C5D2_640.0_640.0.jpg,dynamic/images/2020/03/28/A0950D2037B4C5D2_640.0_640.0.jpg,dynamic/images/2020/03/28/A0950D2037B4C5D2_640.0_640.0.jpg"
	images := strings.Split(urlSring, ",")
	fmt.Println([]rune(urlSring))
	for _, v := range images {
		fmt.Println(v)
	}
}
