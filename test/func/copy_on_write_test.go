package _func

import (
	"fmt"
	"testing"
)

func TestCOWString(t *testing.T) {
	str := "Hello, World"
	str2 := str // 这里并没有复制字符串数据，只是创建了一个新的引
	fmt.Println(str, str2)
	str = "111"
	fmt.Println(str, str2)
}
