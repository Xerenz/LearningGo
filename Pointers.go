package main

import "fmt"

// pointers

func main() {
	type Holder struct {
		foo int
	}

	var ms Holder = Holder{foo: 2}
	// pointer initialization
	var p *Holder = &ms

	// verbose
	ns := Holder{foo: 5}
	n := &ns

	// using new
	var as *Holder
	as = new(Holder)

	(*as).foo = 42 // derefer and modify
	as.foo = 27    // compiler help!

	fmt.Println(p, n, as)

	// slices are projections of arrays
	// slices contains pointers to an underlying array

	var s []int = []int{1, 2, 3} // slice
	x := s

	x[1] = 42

	fmt.Println(x, s)
}
