// Take an existing design and create a copy of it. The copy is then modified to suit a new purpose.
// Copying points to a new object is a costly operation. The prototype pattern is used to avoid this costly operation.
// The prototype pattern is used to create a new object by copying an existing object, known as the prototype.

package main

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	// Shallow copying. The address of both john and jane is the same.
	john := Person{"John", &Address{"123 London Rd", "London", "UK"}}
	jane := john
	jane.Name = "Jane"
	jane.Address.StreetAddress = "124 London Rd"
	println(john.Name, john.Address.StreetAddress)
	println(jane.Name, jane.Address.StreetAddress)

	// Output:
	// John 124 London Rd
	// Jane 124 London Rd

	// Deep copying. The address of both john and jane is different.
	john = Person{"John", &Address{"123 London Rd", "London", "UK"}}
	jane = john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}

	jane.Name = "Jane"
	jane.Address.StreetAddress = "124 London Rd"
	println(john.Name, john.Address.StreetAddress)
	println(jane.Name, jane.Address.StreetAddress)

	// Output:
	// John 123 London Rd
	// Jane 124 London Rd

}
