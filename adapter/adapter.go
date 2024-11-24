package main

import (
	"fmt"
	"strings"
)

type Line struct {
	x1, y1, x2, y2 int
}

type VectorImage struct {
	lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}

// ^^^^ Interface you are given

// vvvv Interface you have
type Point struct {
	x, y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	points := owner.GetPoints()
	buffer := strings.Builder{}
	for _, point := range points {
		buffer.WriteString(fmt.Sprintf("%#v", point))
	}
	return buffer.String()
}

// Solution

type vectorToRasterAdapter struct {
	points []Point
}

func (v *vectorToRasterAdapter) addLine(line Line) {
	left := min(line.x1, line.x2)
	right := max(line.x1, line.x2)
	top := min(line.y1, line.y2)
	bottom := max(line.y1, line.y2)

	if left == right {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if top == bottom {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (v *vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.lines {
		adapter.addLine(line)
	}
	return &adapter
}

func main() {
	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)
	fmt.Print(DrawPoints(a))
}
