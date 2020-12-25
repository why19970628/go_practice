package main

import "fmt"

func main() {
	s1 := make([] int ,3,5)
	s1[0]=1
	s1[1]=2
	fmt.Println("a",s1)
	test(s1)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println("------")

	fmt.Println(s1[2]) // 0
	s2 :=s1[:4] //0123
	fmt.Println(s2)

	s3 :=s1[:3] //0123
	fmt.Println(s3)
	fmt.Println("------")
	main2()
}


func test( s [] int ) {
	s = append(s,3)
	fmt.Println(s)
}



func main2() {
	persons:=make(map[string]int)
	persons["张三"]=19
	mp:=&persons
	fmt.Printf("原始map的内存地址是：%p\n",mp)
	modify(persons)
	fmt.Println("map值被修改了，新值为:",persons)
}

func modify(p map[string]int){
	fmt.Printf("函数里接收到map的内存地址是：%p\n",&p)
	p["张三"]=20
}