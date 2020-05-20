package main

import "fmt"

// collection types

func main() {

	// initializing arrays

	var grades = [3]int{89, 99, 88} // static array
	fmt.Printf("Student grades %v\n", grades)

	subjects := [...]string{"Math", "Physics", "Chemistry"} // dynamic length array
	fmt.Printf("Subjects %v\n", subjects)

	var students [3]string // creating an empty array
	students[0] = "Martin"
	fmt.Printf("Students %v\n", students)

	// length of an array
	fmt.Printf("No. of subjects %v\n", len(subjects))

	// matrix
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[0] = [3]int{0, 1, 0}
	identityMatrix[0] = [3]int{0, 0, 1}

	nullMatrix := [3][3]int{[3]int{0, 0, 0}, [3]int{0, 0, 0}, [3]int{0, 0, 0}}
	fmt.Println(nullMatrix)

	// copying array
	// Go makes full copy of the array

	new := grades
	new[0] = 0 // both grades and new change values

	// Slices

	var slice = []int{1, 2, 3}
	// slice is a dynamic length array

	fmt.Println(slice, len(slice), cap(slice))

	// slice point to the same data
	dum := slice
	dum[0] = 5 // both values change

	// slicing in slice
	newSlice := []int{1, 2, 4, 5, 6, 7, 8, 9}
	a := newSlice[:]  // copy all elements
	b := newSlice[3:] // from 3rd to last
	c := newSlice[:6] // till the 5th element
	d := newSlice[3:6]

	fmt.Println(a, b, c, d)

	// making slices and arrays using make()

	x := make([]int, 3) // first argument is object type
	// second is length
	y := make([]int, 4, 100) // third argument is capacity

	x = append(x, 1, 2, 3)          // apend takes an array, makes copy
	y = append(y, 1, 2, 3, 4, 5, 6) // and adds element, not economical

	fmt.Printf("length : %v, capacity : %v\n", len(x), cap(x))
	fmt.Printf("length : %v, capacity : %v\n", len(y), cap(y))

	// removing elements from a slice

	a = a[1:] // remove first element
	fmt.Println(a)
	a = a[:len(a)-1] // last element
	fmt.Println(a)
	a = append(a[:2], a[3:]...) // removing a middle element
	fmt.Println(a)
}
