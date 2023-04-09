package factory_partten

type Book struct {
	Id   int
	Name string
}

func (b *Book) GetInfo() string {
	return "book info"
}

type Briefs struct {
	Id   int
	Name string
}

func (b *Briefs) GetInfo() string {
	return "briefs info"
}
