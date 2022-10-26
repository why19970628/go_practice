package main

import (
	"errors"
	"log"
)

func Add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case string:
		return a.(string) + b.(string)
	case float64:
		return a.(float64) + b.(float64)
	}
	return errors.New("type not find")
}

func main() {
	log.Print(Add(5, 5))
}
