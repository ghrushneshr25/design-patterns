package main

import "fmt"

// High level module should not depend on Low Level Model
// Both should be dependent on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationShip Relationship
	to           *Person
}

// Low level Module
type RelationsShips struct {
	relations []Info
}

func (r *RelationsShips) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{
		parent, Parent, child,
	})

	r.relations = append(r.relations, Info{
		child, Child, parent,
	})
}

// high level module
type Research struct {
	// Breaking Dependency Inversion principle
	// relationships RelationsShips // this is wrong
	browser RelationshipBrowser
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Research) Investigate() {
	// relations := r.relationships.relations

	// for _, rel := range relations {
	// 	if rel.from.name == "John" && rel.relationShip == Parent {
	// 		fmt.Println("John has a child callled ", rel.to.name)
	// 	}
	// }
	
	result := r.browser.FindAllChildrenOf("John")
	for _, rel := range result {
		fmt.Println("John has a child callled ", rel.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Mary"}
	child2 := Person{"Marlo"}

	relationships := RelationsShips{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}

func (r *RelationsShips) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationShip == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}
