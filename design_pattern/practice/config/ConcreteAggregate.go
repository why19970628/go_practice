package config

type ConcreteAggregate struct {
	eleList []string
}

func (this *ConcreteAggregate) Iterator() Iterator {
	return NewConcreteIterator(this.eleList)
}

func (this *ConcreteAggregate) AddElement(info string) {
	this.eleList = append(this.eleList, info)
}

func NewConcreteAggregate() ConcreteAggregate {
	return ConcreteAggregate{
		eleList: []string{},
	}
}
