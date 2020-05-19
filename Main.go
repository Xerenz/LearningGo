package main

import (
	"fmt"
	"strconv"
)

// declaring var at pkg level

var (
	o    float32
	p    float32 = 66
	name string  = "MyName"

//	Global string  = "MyGlobalString" // variable exposed at global level
)

func main() {
	fmt.Println("Hello Go!")

	// block variables

	var i int // when you declare a variable in main scope
	// and have to use in another local scope
	i = 42

	var j float32 = 27 // when extra info for var

	k := 17

	var l float32
	l = float32(i) // type conversion

	var u string
	u = string(k) // will get unicode

	var a string
	a = strconv.Itoa(k) // convert to ascii string

	fmt.Println(i, j, k, l, u, a)
}
