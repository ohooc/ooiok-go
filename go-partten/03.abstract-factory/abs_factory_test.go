package abstract_factory

import "testing"

func TestFactoryProducer_GetFactory(t *testing.T) {
	producer := NewFactoryProducer()
	shapeFactory := producer.GetFactory("shape")
	colorFactory := producer.GetFactory("color")

	square := shapeFactory.GetShape("square")
	circle := shapeFactory.GetShape("circle")
	square.Draw()
	circle.Draw()

	red := colorFactory.GetColor("red")
	green := colorFactory.GetColor("green")
	red.Fill()
	green.Fill()

}
