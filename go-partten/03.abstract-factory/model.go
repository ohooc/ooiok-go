package abstract_factory

import "fmt"

type Shape interface {
	Draw()
}

type Color interface {
	Fill()
}

type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("画圆")
}

type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("画正方形")
}

type Red struct {
}

func (r *Red) Fill() {
	fmt.Println("填充红色")
}

type Green struct {
}

func (g *Green) Fill() {
	fmt.Println("填充绿色")
}
