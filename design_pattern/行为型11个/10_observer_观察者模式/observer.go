package observer

import "fmt"

type Subject struct {
	arr []*obj
	ctx string
}

func NewNewSubject() *Subject {
	return &Subject{arr: make([]*obj, 0)}
}

func (s *Subject) add(obj2 *obj) {
	s.arr = append(s.arr, obj2)
}

func (s *Subject) notify() {
	for _, v := range s.arr {
		v.update(s)
	}
}

func (s *Subject) update(c string) {
	s.ctx = c
	s.notify()
}

type obj struct {
	name string
}

func NewObj(name string) *obj {
	return &obj{name: name}
}

func (o *obj) update(s *Subject) {
	fmt.Println("obj update: ", o.name, s.ctx)
}
