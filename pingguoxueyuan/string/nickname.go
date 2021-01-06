package main

import (
	"fmt"
	"strings"
)

func NickNameHandler(str string, is_nickname bool) string {
	nickname := ""
	// 昵称
	if is_nickname {
		slice := strings.Split(str, "")
		fmt.Println("slice:", slice, len(slice))
		if len(slice) == 1 {
			nickname = str
		} else if len(slice) == 2 {
			nickname = slice[0] + "*"
		} else if len(slice) >= 3 {
			nickname = slice[0] + "***" + slice[len(slice)-1]
		}
	} else {
		// 手机号 134****5678
		slice := strings.Split(str, "")
		nickname = strings.Join(slice[0:3], "") + "****" + strings.Join(slice[7:], "")
		//nickname = str[0:3] + "****" + str[7:11]
	}
	return nickname
}

func main() {
	//name := NickNameHandler("wang", true)
	fmt.Println(NickNameHandler("王", true))
	fmt.Println(NickNameHandler("王子", true))
	fmt.Println(NickNameHandler("王子厉害", true))
	fmt.Println(NickNameHandler("王%%%", true))
	fmt.Println(NickNameHandler("王.,", true))

	fmt.Println(NickNameHandler("13123107098", false))

}
