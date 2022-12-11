package simple_factory

import "testing"

//TestType1 test get hiapi with factory
func TestType1(t *testing.T) {
	api := NewAnimalSimpleFactory("tiger")
	api.Say()
}

//func TestType2(t *testing.T) {
//	api := NewAPI(2)
//	s := api.Say("Tom")
//	if s != "Hello, Tom" {
//		t.Fatal("Type2 test fail")
//	}
//}
