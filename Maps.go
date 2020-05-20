package main

import "fmt"

// collection types

func main() {

	// initializing maps

	// with entries

	myMap := map[string]int{
		"Martin": 1,
		"Rahul":  4,
		"George": 6,
	}

	// without entries

	newMap := make(map[int]string)

	// array as key map

	arrayAsKey := make(map[[2]string]string)
	arrayAsKey = map[[2]string]string{
		[2]string{"North", "East"}: "NorthEast",
		[2]string{"North", "West"}: "NorthWest",
		[2]string{"South", "East"}: "SouthEast",
		[2]string{"South", "West"}: "SouthWest",
	}

	fmt.Println(myMap, newMap, arrayAsKey)

	// adding to a map
	newMap[1] = "Testing"
	// deleting
	delete(newMap, 1)

	// interogating key in a map

	val, ok := myMap["Martin"]       // ok = true
	val, ok = myMap["SomethingElse"] // ok = false

	fmt.Println(val, ok)

	// structs
	type Pilot struct {
		name   string
		age    int
		planes []string
	}

	aPilot := Pilot{
		name: "John Young",
		age:  20,
		planes: []string{
			"F-22 Raptor",
			"Dasault Rafale",
			"F-35 lightning",
			"Sukhoi Su-30 MKI",
		},
	}

	fmt.Println(aPilot)

	// anonymous structs
	aPlane := struct{ name string }{name: "Mig-21"}

	fmt.Println(aPlane)

	newPlane := aPlane // makes new copy
	newPlane.name = "Mirage-2000"
	fmt.Println(newPlane)

	// struct embedding

	type Plane struct {
		name       string
		wingSpan   float32
		engineType string
		capacity   int
	}

	// FighterPlane inherits from Plane struct
	type FighterPlane struct {
		Plane
		manufacturer string
	}

	fighterPlane := FighterPlane{}

	fighterPlane.name = "Mig-29"
	fighterPlane.wingSpan = 10
	fighterPlane.engineType = "Twin Thrust"
	fighterPlane.capacity = 1
	fighterPlane.manufacturer = "Mikoyan Gerovic"
}
