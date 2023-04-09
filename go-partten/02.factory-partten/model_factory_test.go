package factory_partten

import "testing"

func TestCreateProduct(t *testing.T) {

	info := new(TechFactory).CreateProduct(ProductTechBook).GetInfo()
	t.Log(info)
}
