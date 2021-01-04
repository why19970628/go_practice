package main

import "fmt"

func main() {
	//intChan := make(chan int, 6)
	//
	//for i := 0; i < 6; i++ {
	//	intChan <- i
	//}
	//
	//stringChan := make(chan string, 6)
	//for i := 0; i < 6; i++ {
	//	stringChan <- "data:" + fmt.Sprintf("%d", i)
	//}
	//
	//for {
	//	select {
	//	case val := <-intChan:
	//		fmt.Println(val)
	//	case val := <-stringChan:
	//		fmt.Println(val)
	//	default:
	//		//fmt.Println("做其他事")
	//		return
	//	}
	//}

	fmt.Println("------")
	main1()

	fmt.Println("22222")
	main2()

}

func f1(in chan int) {
	fmt.Println(<-in)
}

func main1() {
	out := make(chan int)
	go f1(out) // 必须先消费
	out <- 2
}

func main2() {
	out := make(chan int, 1) //buffer大于0
	out <- 2
	go f1(out)
}
