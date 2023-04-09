package abstract_factory

type AbstractFactory interface {
	GetShape(shape string) Shape
	GetColor(color string) Color
}

type ShapeFactory struct {
}

func (s *ShapeFactory) GetShape(shape string) Shape {
	switch shape {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		return nil
	}
}

func (s *ShapeFactory) GetColor(color string) Color {
	return nil
}

type ColorFactory struct {
}

func (c *ColorFactory) GetShape(shape string) Shape {
	return nil
}

func (c *ColorFactory) GetColor(color string) Color {
	switch color {
	case "red":
		return &Red{}
	case "green":
		return &Green{}
	default:
		return nil
	}
}

type FactoryProducer struct {
}

func NewFactoryProducer() *FactoryProducer {
	return &FactoryProducer{}
}

func (f *FactoryProducer) GetFactory(name string) AbstractFactory {
	switch name {
	case "color":
		return &ColorFactory{}
	case "shape":
		return &ShapeFactory{}
	default:
		return nil
	}
}
