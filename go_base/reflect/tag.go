package main

import (
	"fmt"
	"reflect"
)

func main() {
	var u User
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag.Get("json"), ",", sf.Tag.Get("bson"))
	}

}

type User struct {
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_age"`
}
