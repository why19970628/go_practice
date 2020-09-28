package main

import "fmt"

func testArray1()  {
	var a [5]int
	a[0] = 200
	a[2] = 100
	fmt.Println(a)
}
func testArray2()  {
	var a [3]int = [3]int{100,200,300}
	//a[0] = 200
	//a[2] = 100
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("array[%d] is %d\n",i, a[i])
	}
	//var index, value int
	for index, value := range a {
		fmt.Println("for range:", index, value)
	}
}

func testArray3()  {
	var a []int
	//a[0] = 200
	//a[2] = 100
	fmt.Println(a)
}

func testArray4()  {
	a:= [5]int{2:100,4:400}
	//a[0] = 200
	//a[2] = 100
	fmt.Println(a)
}


func main()  {
	testArray1()
	testArray2()
	testArray3()
	testArray4()

}
