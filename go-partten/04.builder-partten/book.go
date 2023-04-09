package builder_partten

type Book struct {
	Id    int
	Name  string
	Price float64
}

func (b *Book) Builder(id int, name string) *BookBuilder {
	return NewBookBuilder(id, name)
}
