package main

import "fmt"

type User1 struct {
	Name string
	Id int
}

func main()  {
	user := &User1{Name: "wang", Id: 1}
	fmt.Println(user)
	user.Name = "hua"
	fmt.Println(user)

}


