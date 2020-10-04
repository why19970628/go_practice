package main

import "fmt"

type Vertex struct {

	X, Y int

}

func main() {

	 v := new(Vertex)

	 fmt.Println(v)

	v.X, v.Y = 11, 9

	fmt.Println(v)

}

//func main() {
//	var i *int
//	i=new(int)
//	fmt.Println(i)
//	*i=10
//	fmt.Println(*i)
//
//	var b string
//	b = "666"
//	fmt.Println(b)
//}