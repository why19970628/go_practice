package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func p2() {
	x := []int{1, 2, 3}
	y := x[1:]
	fmt.Printf("222 x: %v %p, len:%v,cap:%v \n", x, &x, len(x), cap(x)) // [1 2 3]  len:3,cap:3
	fmt.Printf("222 y: %v %p, len:%v,cap:%v \n", y, &y, len(y), cap(y)) // [2 3] len:2,cap:2
	y[0] = 10
	fmt.Printf("333 y: %v %p, len:%v,cap:%v \n", y, &y, len(y), cap(y)) // [10 3] len:2,cap:2
	fmt.Printf("333 x: %v %p, len:%v,cap:%v \n", x, &x, len(x), cap(x)) // [1 10 3]   len:3,cap:3

	y = append(y, 50)
	fmt.Printf("444 y: %v %p, len:%v,cap:%v \n", y, &y, len(y), cap(y)) // [10 3 50] len:3,cap:4
	fmt.Printf("444 x: %v %p, len:%v,cap:%v \n", x, &x, len(x), cap(x)) // [1 10 3]   len:3,cap:3

}

func p1() {
	x := []int{1, 2, 3}
	y := x[:2]
	fmt.Printf("222 x: %v %p, len:%v,cap:%v \n", x, &x, len(x), cap(x)) // [1 2 3]  len:3,cap:3
	fmt.Printf("222 y: %v %p, len:%v,cap:%v \n", y, &y, len(y), cap(y)) // [1 2] len:2,cap:3
	y = append(y, 50)
	fmt.Printf("333 x: %v %p, len:%v,cap:%v \n", x, &x, len(x), cap(x)) // [1 2 50]  len:3,cap:3
	fmt.Printf("333 y: %v %p, len:%v,cap:%v \n", y, &y, len(y), cap(y)) // [1 2 50]  len:3,cap:3

	fmt.Println("-------------------")
	y = append(y, 60)            // 超出容量,公用底层的切片会重新分配一个地址 x y 此时独立
	fmt.Printf("%v %p\n", y, &y) //  output: [1 2 50 60]
	y[0] = 20
	fmt.Printf("%v %p\n", y, &y)        // output:  [20 2 50 60]
	fmt.Printf("555 x: %v %p\n", x, &x) // [1 2 50] 0xc00000c060
}

func main() {
	p2()
}

type Slice []int

func (A Slice) Append(value int) {
	A1 := append(A, value)
	fmt.Printf("%p\n%p\n", A, A1) // 地址一样

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&A))
	fmt.Printf("A Data:%d,Len:%d,Cap:%d\n", sh.Data, sh.Len, sh.Cap)

	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&A1))
	fmt.Printf("A1 Data:%d,Len:%d,Cap:%d\n", sh1.Data, sh1.Len, sh1.Cap)

}

func append_2() {
	mSlice := make(Slice, 10, 20)
	mSlice.Append(5)
	fmt.Println(mSlice)
}

// 通过a[i:j]生成的切片，两个切片底层竟然指定同一个数组。
// 导致在没有超过cap的情况下，对第二个切片操作会影响第一个切片。
func append_3() {
	arr := []int{1, 2, 3, 4, 5, 6}
	sh0 := (*reflect.SliceHeader)(unsafe.Pointer(&arr))
	fmt.Printf("sh0 Data:%d,Len:%d,Cap:%d\n", sh0.Data, sh0.Len, sh0.Cap) // sh0 Data:824633811232,Len:6,Cap:6

	slice := arr[1:2] // data num:2 len 1 cap:5
	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("A1 Data:%d,Len:%d,Cap:%d\n", sh1.Data, sh1.Len, sh1.Cap) // A1 Data:824633811240,Len:1,Cap:5

	// 不会发生扩容
	slice = append(slice, 7, 8, 9)
	fmt.Println(slice) // 2 7 8 9
	fmt.Println(arr)   // 127896
}
