package main

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
)

func JsonToMap() {
	jsonStr := `
    {
        "name":"liangyongxing",
        "age":12
    }
    `
	var mapResult map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Println(err)
	}
	fmt.Println(mapResult)
}

type Person struct {
	Name string
	Age  int
}

func MapToStruct() {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang637210"
	mapInstance["Age"] = 28
	fmt.Println(mapInstance)

	var person Person
	//将 map 转换为指定的结构体
	if err := mapstructure.Decode(mapInstance, &person); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("map2struct后得到的 struct_addr 内容为:%v", person)
}

func main() {
	//JsonToMap()
	MapToStruct()
}
