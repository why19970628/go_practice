package main

func firstNonRepeatingChar(s string) string {
	charCount := make(map[rune]int64)
	for _, v := range s {
		charCount[v]++
	}
	for _, c := range s {
		if charCount[c] == 1 {
			return string(c)
		}
	}

	return " "
}
