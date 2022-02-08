package singleton

import (
	"testing"
)

//TestType1 test get hiapi with factory
func TestSingleton(t *testing.T) {
	instance := GetInstance()
	instance.foo()

	//ins1 := GetInstance()
	//ins2 := GetInstance()
	//if ins1 != ins2 {
	//	t.Fatal("instance is not equal")
	//}
	//s := single()
	//s.t()
}
