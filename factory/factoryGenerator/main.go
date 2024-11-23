package main

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)

	developer := developerFactory("adam")
	manager := managerFactory("jane")

	println(developer)
	println(manager)

	/*
		Approach 2
		Can create multiple factories with different configurations
		We can modify the configuration of the factory after creating it
	*/
	bossFactory := NewEmployeeFactory2("CEO", 100000)
	bossFactory.AnnualIncome = 200000
	boss := bossFactory.Create("Sam")
	println(boss)

}
