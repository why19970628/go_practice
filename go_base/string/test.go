package main

import "fmt"

func main() {
	fieldNums := 28
	quesMarkString := "("
	for i := 0; i < fieldNums; i++ {
		quesMarkString += "?, "
	}
	quesMarkString = quesMarkString[:len(quesMarkString)-2] + ")"
	fmt.Println(quesMarkString, len(quesMarkString))

	//valueStrings := make([]string, 0, len(orders))

	fmt.Println(fmt.Sprintf("{\"product\":\"%s\"}", "333"))

}
