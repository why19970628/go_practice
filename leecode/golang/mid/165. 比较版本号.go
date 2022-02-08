package main

import (
	"fmt"
	"strconv"
	"strings"
)

//func compareVersion(version1, version2 string) int {
//	v1 := strings.Split(version1, ".")
//	v2 := strings.Split(version2, ".")
//	for i := 0; i < len(v1) || i < len(v2); i++ {
//		x, y := 0, 0
//		if i < len(v1) {
//			x, _ = strconv.Atoi(v1[i])
//		}
//		if i < len(v2) {
//			y, _ = strconv.Atoi(v2[i])
//		}
//		if x > y {
//			return 1
//		}
//		if x < y {
//			return -1
//		}
//	}
//	return 0
//}

func compareVersion2(version1, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")

	maxLength := 0
	if len(s1) > len(s2) {
		maxLength = len(s1)
	} else {
		maxLength = len(s2)
	}
	for i := 0; i < maxLength; i++ {
		a, b := 0, 0
		if len(s1) > i {
			a, _ = strconv.Atoi(s1[i])
		}
		if len(s2) > i {
			b, _ = strconv.Atoi(s2[i])
		}

		if a > b {
			return 1
		}
		if a < b {
			return -1
		}
	}
	return 0
}

func main() {
	var version1, version2 string = "1.01", "1.001"
	fmt.Println(compareVersion2(version1, version2))
}
