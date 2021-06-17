package main

import "fmt"

func testArray1() {
	var a [5]int
	a[0] = 200
	a[2] = 100
	fmt.Println(a)
}
func testArray2() {
	var a [3]int = [3]int{100, 200, 300}
	//a[0] = 200
	//a[2] = 100
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("array[%d] is %d\n", i, a[i])
	}
	//var index, value int
	for index, value := range a {
		fmt.Println("for range:", index, value)
	}
}

func testArray3() {
	var a []int
	//a[0] = 200
	//a[2] = 100
	fmt.Println(a)
}

func testArray4() {
	a := [5]int{2: 100, 4: 400}
	//a[0] = 200
	//a[2] = 100
	fmt.Println(a)
}

func array_pop() {
	array := []int{1, 2, 3, 4}
	fmt.Println(array)
	//删除第i个元素
	i := 2
	array = append(array[:i], array[i+1:]...)
	fmt.Println(array)
}

func array() {
	slice := make([]int, 5) // 一次性拿到连续内存空间的切片，用指针访问会更好一些
	//slice := make([]int, 0, 5) // 一次性拿到连续内存空间
	slice = append(slice, 1)
	fmt.Println(slice)

	slice2 := make([]int, 0, 5) // 一次性拿到连续内存空间的切片，用指针访问会更好一些
	slice2 = append(slice2, 1)
	fmt.Println(slice2)
}

func main() {
	//testArray1()
	//testArray2()
	//testArray3()
	//testArray4()
	//array_pop()
	array()
}
