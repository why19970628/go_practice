package model

//import "fmt"
//
///***
// *  Simple Factory
// */
//type FoodFactory struct {
//}
//
//func (ff FoodFactory) CreateFood(name string) Food {
//	var s Food;
//	switch name {
//	case "Meat":
//		s = new(Meat);
//	case "Hamberger":
//		s = new(Hamberger)
//	}
//	return s;
//}
//
//type Food interface {
//	Eat()
//}
//
//type Meat struct {
//}
//
//type Hamberger struct {
//}
//
//func (m Meat) Eat() {
//	fmt.Println("Eat meat.")
//}
//
//func (h Hamberger) Eat() {
//	fmt.Println("Eat Hamberger.")
//}
//
//func main2(){
//
//	// Simple Factory
//	f := FoodFactory{}
//	f.CreateFood("Meat").Eat()
//	f.CreateFood("Hamberger").Eat()
//}