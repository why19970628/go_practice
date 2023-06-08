package backtrack

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func restoreIpAddresses(s string) []string {
	pathStr, resStr = make([]string, 0), make([]string, 0)
	restoreIpAddressesDfs(s, 0)
	return resStr
}

func restoreIpAddressesDfs(s string, start int) {
	if len(pathStr) == 4 {
		// 指向最后
		if start == len(s) {
			str := strings.Join(pathStr, ".")
			resStr = append(resStr, str)
		}
		return
	}
	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			pathStr = append(pathStr, str) // 符合条件的就进入下一层,继续匹配
			restoreIpAddressesDfs(s, i+1)
			pathStr = pathStr[:len(pathStr)-1]
		} else {
			break
		}
	}
}

func TestRestoreIpAddresses(t *testing.T) {
	fmt.Println(restoreIpAddresses("0000"))
}
