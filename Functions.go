package main

import "fmt"

func main() {
	name := "Martin"
	greeting := "Hello"
	passByValue(greeting, name)
	fmt.Println(greeting, name)
	passByRefer(&greeting, &name)
	fmt.Println(greeting, name)

	sum(67, 88, 56, 1, 3, 56)

	res := max(23, 11, 45, 2, 67)

	fmt.Println(*res)

	res, err := devide(2.0, 0.0) // normal flow of code
	if err != nil {
		fmt.Println(err) // only exception indented
		return           // end of execution
	}
	fmt.Println(res) // remains left indented

	g := Greet{name: "Martin", msg: "Hello"}
	g.greet()
	g.change("Goodbye")
	g.greet()
}

// variant parameters
func sum(values ...int) {
	res := 0
	for _, v := range values {
		res += v
	}

	fmt.Println(res)
}

func sugarSum(values ...int) (res int) { // res init as 0
	for _, v := range values {
		res += v
	}
	return // returns res
}

func mul(msg string, values ...int) int { // variant parameter should always be last
	res := 1
	for _, val := range values {
		res *= val
	}

	return res // returns copy of variable
}

func max(values ...int) *int {

	maxVal := 0

	for _, v := range values {
		if maxVal < v {
			maxVal = v
		}
	}
	// for such func memory moves from
	// local stack to the common heap
	return &maxVal // so this val persists
}

func passByValue(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Name Changed"
}

func passByRefer(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Reference"
}

// multiple return
func devide(num, deno float64) (float64, error) {
	if deno == 0.0 {
		return 0.0, fmt.Errorf("Cannot devide by Zero")
	}
	return num / deno, nil
}

// methods
type Greet struct {
	name string
	msg  string
}

func (g Greet) greet() {
	// this is a method
	fmt.Println(g.msg, g.name)
}

// pointer method
func (g *Greet) change(newMsg string) {
	g.msg = newMsg
}
