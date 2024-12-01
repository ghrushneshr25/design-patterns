package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("<->", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}

}

func NewCircle(Color string) *GraphicObject {
	return &GraphicObject{"Circle", Color, nil}
}

func NewSquare(Color string) *GraphicObject {
	return &GraphicObject{"Square", Color, nil}
}

func main() {
	drawings := GraphicObject{"My Drawing", "", nil}

	drawings.Children = append(drawings.Children, *NewCircle("Red"))
	drawings.Children = append(drawings.Children, *NewSquare("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Blue"))

	drawings.Children = append(drawings.Children, group)
	fmt.Println(drawings.String())
}
