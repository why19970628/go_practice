package main

import "fmt"

func reverse(testString string)  {
	runes := []rune(testString)
	fmt.Println(runes)
	data := ""
	for i := 0; i < len(runes); i++ {
		data = data + string(runes[len(runes)-i-1])
	}

	fmt.Println(data)

	str := "hello 世界"
	fmt.Println("rune:", len([]rune(str)))
}
func main() {
	testString := "hello 世界-01"

	reverse(testString)
}
