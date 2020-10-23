package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "how do you do"
	wordCount := make(map[string]int, 10)

	words := strings.Split(s, " ")
	fmt.Println(words)

	for _, value := range words {
		fmt.Println(value)
		v, ok := wordCount[value]
		if ok{
			wordCount[value] = v + 1
		}else {
			wordCount[value] =  1

		}
	}
	fmt.Println(wordCount)


}
