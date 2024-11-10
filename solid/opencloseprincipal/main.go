package main

import "fmt"

// OCP
// Open for extension, closed for modification
// specification

type Color int

const (
	RED Color = iota
	GREEN
	BLUE
)

type Size int

const (
	SMALL Size = iota
	MEDIUM
	LARGE
)

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {

	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {

	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterByColorAndSize(products []Product, size Size, color Color) []*Product {

	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

type Product struct {
	name  string
	color Color
	size  Size
}

// BASED ON OPEN CLOSED PRINCIPAL

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (c SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == c.size
}

type AndSpecificaton struct {
	first, second Specification
}

func (a AndSpecificaton) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(product []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range product {
		if spec.IsSatisfied(&v) {
			result = append(result, &product[i])
		}
	}
	return result
}

func main() {
	apple := Product{
		"Apple", GREEN, SMALL,
	}

	tree := Product{
		"Tree", GREEN, LARGE,
	}

	house := Product{
		"House", BLUE, LARGE,
	}

	products := []Product{apple, tree, house}
	fmt.Printf("Green Products : \n")
	f := Filter{}

	for _, v := range f.FilterByColor(products, GREEN) {
		fmt.Printf(" - %s is green \n", v.name)
	}

	fmt.Println("\nGree Products based on New Implementation")
	greenSpec := ColorSpecification{GREEN}

	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Println("\nLarge and Green Products based on New Implementation")
	largeSpec := SizeSpecification{LARGE}
	lgSpec := AndSpecificaton{
		largeSpec, greenSpec,
	}

	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is green and large\n", v.name)
	}
}
