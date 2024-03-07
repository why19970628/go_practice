package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	n1 := len(word1)
	n2 := len(word2)
	newStr := []byte{}
	minN := min(n1, n2)
	for i := 0; i < minN; i++ {
		newStr = append(newStr, word1[i])
		newStr = append(newStr, word2[i])
	}
	newStr = append(newStr, word1[minN:]...)
	newStr = append(newStr, word2[minN:]...)
	return string(newStr)

}

func main() {

	// ap bq cd
	fmt.Println(mergeAlternately("abcd", "pq"))
	//fmt.Println(mergeAlternately("abc", "pqr"))
}
