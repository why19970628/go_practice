package main

import (
	"fmt"
	"time"
)

type OrderHandler interface {
	run()
	getLogId() int
}

type SaveOrder struct {
	desc string
	id   int
}

func NewSaveOrder(desc string, id int) *SaveOrder {
	return &SaveOrder{desc: desc, id: id}
}

func (s *SaveOrder) run() {
	fmt.Println(s.id)
}

func (s *SaveOrder) saveLog() {
}

func (s *SaveOrder) getLogId() int {
	fmt.Println(s.id)
	return s.id
}

type UpdateOrder struct {
	orderID int
}

func NewUpdateOrder(orderID int) *UpdateOrder {
	return &UpdateOrder{orderID: orderID}
}

func (s *UpdateOrder) run() {
	fmt.Println(s.orderID)
}

func (s *UpdateOrder) getLogId() int {
	return s.orderID
}

type f func()

func saveLog(run OrderHandler) f {
	return func() {
		go func() {
			fmt.Println("save:", run.getLogId())
		}()
		fmt.Println("start")
		run.run()
		fmt.Println("end")
	}
}

func main() {
	order := NewSaveOrder("保存订单", 123)
	saveLog(order)()
	time.Sleep(time.Second)

	order2 := NewUpdateOrder(456)
	saveLog(order2)()
	time.Sleep(time.Second)
}
