package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"io/ioutil"
	"log"
	// "net/http"
)

// Defer
// defer executes the statement just before the function returns
// defer follows lifo execution

func main() {

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // defer allows opening and closing to be together

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)

	// defer takes the value of variable before its called

	a := "starts"
	defer fmt.Println(a) // prints start
	a = "end"

	fmt.Println("start")
	defer fmt.Println("defer happened") // defer happens before panic
	panic("panic happened")
	fmt.Println("end")

	// we can use this quality of defer to handle panic

	fmt.Println("This is the main exec func")
	panicker() // func with error
	fmt.Println("Normal execution")
	// the funcs higher up the call stack will always
	// continue their exec
	// but func which raise err stop
}

// a function which would fail
func panicker() {
	fmt.Println("This function is about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// panic(err) // rethrow panic if it cannot be dealt with
		}
	}()
	panic("Something bad happened")
	fmt.Println("End") // func exec stops so no print
	// func exit and go back to main
}
