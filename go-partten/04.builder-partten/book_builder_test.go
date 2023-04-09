package builder_partten

import "testing"

func TestBuildBook(t *testing.T) {
	book := new(Book).Builder(111, "golang").SetPrice(88.8).Build()
	t.Log(book)
}
