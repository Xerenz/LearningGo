package main

import "fmt"

func main() {
	animals := []Animal{Dog{}, Cat{}, Goat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speaks())
	}
}

type Animal interface {
	Speaks() string
}

// whether or not a type implements an interface
// is determined automatically in go

// defining types
type Dog struct{}
type Cat struct{}
type Goat struct{}

// defining methods
func (d Dog) Speaks() string {
	return "Woof!"
}

func (c Cat) Speaks() string {
	return "Meow!"
}

func (g Goat) Speaks() string {
	return "Mehh!"
}
