package main


import "fmt"

func main()  {
	user := &User{Name: "wang", Id: 1}
	fmt.Println(user)
	user.Name = "hua"
	fmt.Println(user)

}


