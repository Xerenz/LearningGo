package main

import "fmt"

func main() {
	// using if statements with initialiser

	aMap := map[int]string{
		1: "Something",
		2: "SomethingElse",
	}

	if pop, ok := aMap[1]; ok {
		fmt.Println(pop)
	}

	// interfacing and switch

	var i interface{} = 1

	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float32:
		fmt.Println("i is an float")
	case string:
		fmt.Println("i is an string")
	case [3]int:
		fmt.Println("i is an int array")
	default:
		fmt.Println("i is another type")
	}

}
