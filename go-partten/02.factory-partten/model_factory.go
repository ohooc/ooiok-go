package factory_partten

const (
	ProductTechBook = iota
	ProductDailyBriefs
)

type ProductType int

type IProductFactory interface {
	CreateProduct(t ProductType) IProduct
}

type IProduct interface {
	GetInfo() string
}

type TechFactory struct {
}

func (tf *TechFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductTechBook:
		return &Book{}
	}
	return nil
}

type DailyFactory struct {
}

func (df *DailyFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductDailyBriefs:
		return &Briefs{}
	}
	return nil
}
