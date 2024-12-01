package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle with radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius = c.Radius * factor
}

func NewCircle(radius float32) *Circle {
	return &Circle{Radius: radius}
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

func main() {
	c := Circle{2}
	c.Resize(2)
	fmt.Println(c.Render())

	redCircle := ColoredShape{&c, "Red"}
	fmt.Println(redCircle.Render())

	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
