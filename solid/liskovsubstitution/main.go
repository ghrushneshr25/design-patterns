package main

import "fmt"

// Liskov Substitution Principle

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetWidth(w int) {
	r.width = w
}

func (r *Rectangle) SetHeight(h int) {
	r.height = h
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

type Square2 struct {
	size int
}

func NewSquare2(size int) *Square2 {
	sq := Square2{size}
	return &sq
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)

	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()

	fmt.Print("Expected : ", expectedArea, " Actual : ", actualArea, "\n")
}

func main() {
	rec := &Rectangle{2, 3}
	UseIt(rec)

	sq := NewSquare(2)
	UseIt(sq)

}
