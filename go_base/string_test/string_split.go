package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)

	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	fmt.Println(result, reflect.TypeOf(result))
	return string(result)
}

func main() {
	ips := "10.22.11.231"
	ipSlice := strings.Split(ips, ".")
	fmt.Println(GetRandomString(5))

	for index, value := range ipSlice {
		fmt.Printf("ip index = %d, value = %s\n", index, value)
	}
}
