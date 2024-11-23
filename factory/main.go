package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

// This is a factory function
func NewPerson(name string, age int) *Person {
	if age < 16 {
		// can add validation here..
	}
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("John", 30)

	fmt.Println(p)
}
