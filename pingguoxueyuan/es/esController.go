package main

import (
	"context"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func EsInsert() {
	p1 := Person{Name: "rion", Age: 23, Married: false}
	put1, err := EsClient.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
//根据id查询
func Get(){
	get ,err:= EsClient.Get().Index("ecommerce").Type("product").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get.Found {
		fmt.Println(get.Source)
	}
}

func main() {
	fmt.Println("66")
	//EsInsert()
	//Get()
}
