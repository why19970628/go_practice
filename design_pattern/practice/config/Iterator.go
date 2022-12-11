package config

type Iterator interface {
	HashNext() bool
	Next() interface{}
	Remove()
}

type ConcreteIterator struct {
	eleList []string
	offset  int
}

func (this *ConcreteIterator) HashNext() bool {
	if this.offset < len(this.eleList) {
		return true
	}
	return false
}

func (this *ConcreteIterator) Next() interface{} {
	element := this.eleList[this.offset]
	this.offset += 1
	return element
}

func (this *ConcreteIterator) Remove() {
	index := copy(this.eleList[this.offset:], this.eleList[this.offset+1:])
	this.eleList = this.eleList[:this.offset+index]
}

func NewConcreteIterator(eleList []string) *ConcreteIterator {
	return &ConcreteIterator{
		eleList: eleList,
	}
}
