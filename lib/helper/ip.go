package helper

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	fmt.Println(ipSegs)
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		fmt.Println("1", ipSeg, pos, tempInt, ipInt)
		// 相加
		ipInt = ipInt | tempInt
		fmt.Println("2", ipInt)
		pos -= 8
	}
	return ipInt
}

// 4294967295
func StringIpToInt2(ipstring string) int {
	respInt := 0
	pos := 24
	s := strings.Split(ipstring, ".")
	for _, num := range s {
		n, _ := strconv.Atoi(num)
		tempInt := n << pos
		respInt = respInt | tempInt
		pos -= 8
	}
	return respInt
}

func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}
