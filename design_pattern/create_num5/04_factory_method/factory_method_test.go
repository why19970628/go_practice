package factory_method

import (
	"fmt"
	"testing"
)

//func compute(factory OperatorFactory, a, b int) int {
//	op := factory.Create()
//	op.SetA(a)
//	op.SetB(b)
//	return op.Result()
//}
//
func TestOperator(t *testing.T) {
	//	var (
	//		factory OperatorFactory
	//	)
	//
	//	factory = PlusOperatorFactory{}
	//	if compute(factory, 1, 2) != 3 {
	//		t.Fatal("error with factory method pattern")
	//	}
	//
	//	//factory = MinusOperatorFactory{}
	//	//if compute(factory, 4, 2) != 2 {
	//	//	t.Fatal("error with factory method pattern")
	//	//}

	fac := &(AddFactory{})
	oper := fac.CreateOperation()
	oper.SetA(1)
	oper.SetB(2)
	fmt.Println(oper.GetResult())
}
