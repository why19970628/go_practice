package main

type Book struct {
	Id    int
	Name  string
	Price float64
}

func NewBook(id int, name string, price float64) *Book {
	return &Book{Id: id, Name: name, Price: price}
}

func (b *Book) Builder(id int, name string) *BookBuilder {
	return NewBookBuilder(id, name)
}

// ---------------------
type BookBuilder struct {
	id    int
	name  string
	price float64
}

func NewBookBuilder(id int, name string) *BookBuilder {
	return &BookBuilder{id: id, name: name}
}

func (this *BookBuilder) Build() *Book {
	book := &Book{Id: this.id, Name: this.name}
	if this.price > 0 {
		book.Price = this.price
	}
	return book
}

func (this *BookBuilder) SetPrice(price float64) *BookBuilder {
	this.price = price
	return this
}
