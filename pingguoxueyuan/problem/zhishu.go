package main

import (
	"fmt"
)


func justify(n int)  bool{
	if n<=1{
		return false
	}else {
		for i:=2;i<n;i++{
			if n %i == 0{
				return false
			}
		}
		return true
	}
}

func example()  {
	for i:=2;i<100;i++{
		if justify(i) == true{
			fmt.Println(i, "为质数")
		}
	}
}
func main()  {
	example()
}
