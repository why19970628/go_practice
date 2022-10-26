package main

import (
	"fmt"
	"log"
)

type JianshuOrderRes struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    JianshuOrderData `json:"data"`
}

type JianshuOrderData struct {
	Url string `json:"url"`
}

func Print(data interface{}) {
	log.Printf("%+v\n", data)
}
func main() {
	res := JianshuOrderRes{}
	fmt.Println(res)
	Print(res)
}
