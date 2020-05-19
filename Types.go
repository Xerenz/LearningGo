package main

import "fmt"

// Primitive types in Go

func main() {
	// strings are immutable arrays of bytes
	s := "This is my Go programming"
	fmt.Println(s)

	// converting the string to bytes
	b := []byte(s)
	fmt.Println(b)

	fmt.Printf("%v, %T\n", s[2], s[2]) // each element is an int byte

	// Runes
	// Runes are utf-32 types

	r := 'a'                     // runes are declared by single quotes
	fmt.Printf("%v, %T\n", r, r) // runes are essentially int32 types
}
