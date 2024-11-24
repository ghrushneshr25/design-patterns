package main

import "fmt"

// Circle , square
// Raster , vector
// RasterCircle , RasterSquare, VectorCircle, VectorSquare

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Vector drawing a circle of radius ", radius)
}

type RasterRenderer struct {
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Raster drawing a circle of radius ", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	raster := RasterRenderer{}
	// vector := VectorRenderer{}

	circle := NewCircle(&raster, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
}
