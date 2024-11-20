package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (builder *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	builder.person.StreetAddress = streetAddress
	return builder
}

func (builder *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	builder.person.City = city
	return builder
}

func (builder *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	builder.person.Postcode = postcode
	return builder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (builder *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	builder.person.CompanyName = companyName
	return builder
}

func (builder *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	builder.person.Position = position
	return builder
}

func (builder *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	builder.person.AnnualIncome = annualIncome
	return builder
}

func (builder *PersonBuilder) Build() *Person {
	return builder.person
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
			At("123 London Road").
			In("London").
			WithPostcode("SW12BC").
		Works().
			At("Fabrikam").
			AsA("Programmer").
			Earning(123000)

	person := pb.Build()
	fmt.Println(person)
}
