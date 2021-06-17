//demo_7.go
package main

import (
	"fmt"
)

func main() {
	//一维数组
	var arr_1 [5] int
	fmt.Println(arr_1)

	var arr_2 =  [5] int {1, 2, 3, 4, 5}
	fmt.Println(arr_2)

	arr_3 := [5] int {1, 2, 3, 4, 5}
	fmt.Println(arr_3)

	arr_4 := [...] int {1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)

	arr_5 := [5] int {0:3, 1:5, 4:6}
	fmt.Println(arr_5)

	//二维数组
	var arr_6 = [3][5] int {{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	arr_7 :=  [3][5] int {{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	arr_8 :=  [...][5] int {{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0:3, 1:5, 4:6}}
	fmt.Println(arr_8)

	var sli_1 [] int      //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_1),cap(sli_1),sli_1)

	var sli_2 = [] int {} //空切片
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_2),cap(sli_2),sli_2)

	var sli_3 = [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_3),cap(sli_3),sli_3)

	sli_4 := [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_4),cap(sli_4),sli_4)

	var sli_5 [] int = make([] int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_5),cap(sli_5),sli_5)

	sli_6 := make([] int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_6),cap(sli_6),sli_6)
}