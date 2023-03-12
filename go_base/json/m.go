package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	F1 string `json:"f_1"`
	I2 int64  `json:"i_2"`
}

type m interface {
	client(data []byte, resp interface{})
}

type done struct {
}

func (done) client(data []byte, resp interface{}) {
	_ = json.Unmarshal(data, resp)
}

func main() {
	stu := A{
		F1: "aaaa",
		I2: 66666,
	}
	data, err := json.Marshal(stu)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	var a A
	done{}.client(data, &a)
	fmt.Println(a)
	//json.Unmarshal(data)
}

//stu := A{}
//data, err := json.Marshal(&stu)
//if err != nil {
//fmt.Printf("序列化错误 err=%v\n", err)
//return
//	str :=
