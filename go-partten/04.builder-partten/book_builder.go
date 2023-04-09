package builder_partten

type BookBuilder struct {
	id    int
	name  string
	price float64
}

func NewBookBuilder(id int, name string) *BookBuilder {
	return &BookBuilder{id: id, name: name}
}

func (b *BookBuilder) Build() *Book {
	book := &Book{Id: b.id, Name: b.name}
	if b.price > 0 {
		book.Price = b.price
	}
	return book
}

func (b *BookBuilder) SetPrice(price float64) *BookBuilder {
	b.price = price
	return b
}
