package main

import "fmt"

// exposing the interface rather than the struct
type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Println("Hello, my name is", p.name)
}

func NewPerson(name string, age int) Person {
	return &person{name, age}
}

func main() {
	p := NewPerson("John", 30)

	p.SayHello()
}
