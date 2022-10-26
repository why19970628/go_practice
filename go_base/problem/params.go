package main

import "log"

var str []string

func main() {
	setVal(&str)
	log.Println(str)
}

//需要在这里赋值str，但是又不能直接引用 str
func setVal(val *[]string) {
	*val = []string{"a", "b"}
}
