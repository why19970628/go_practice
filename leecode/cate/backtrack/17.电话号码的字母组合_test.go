package backtrack

import (
	"fmt"
	"testing"
)

func letterCombinations(digits string) (res []string) {
	if digits == "" {
		return
	}
	mp = make(map[byte]string)
	mp['2'] = "abc"
	mp['3'] = "def"
	mp['4'] = "ghi"
	mp['5'] = "jkl"
	mp['6'] = "mno"
	mp['7'] = "pqrs"
	mp['8'] = "tuv"
	mp['9'] = "wxyz"
	letterCombinationsDfs(digits, 0)
	return resStr
}

func letterCombinationsDfs(digits string, currIndex int) {
	if len(pathStr) == len(digits) {
		tmp := string(pathBt)
		resStr = append(resStr, tmp)
		return
	}
	// '2'
	digit := digits[currIndex]
	str := mp[digit]
	// "abc"
	for i := 0; i < len(str); i++ {
		pathBt = append(pathBt, str[i])
		letterCombinationsDfs(digits, currIndex+1)
		pathBt = pathBt[:len(pathBt)-1]
	}
}

func letterCombinationsV2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mp = make(map[byte]string)
	arrResp = make([]string, 0)
	mp['2'] = "abc"
	mp['3'] = "def"
	mp['4'] = "ghi"
	mp['5'] = "jkl"
	mp['6'] = "mno"
	mp['7'] = "pqrs"
	mp['8'] = "tuv"
	mp['9'] = "wxyz"
	letterCombinationsV2Dfs(digits, 0)
	return arrResp
}

var (
	arrResp []string
	temp    []byte
)

func letterCombinationsV2Dfs(digits string, start int) {
	if len(temp) > 0 && len(temp) == len(digits) {
		arrResp = append(arrResp, string(temp))
		return
	}
	str := mp[digits[start]]
	// abc
	for j := 0; j < len(str); j++ {
		temp = append(temp, str[j])
		letterCombinationsV2Dfs(digits, start+1)
		temp = temp[:len(temp)-1]
	}
}

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinationsV2("2"))
}
