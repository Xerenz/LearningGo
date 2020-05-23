package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/** use -race flag when running go to check for race condition **/

var wg = sync.WaitGroup{} // define a wait group object
var counter = 0
var mutex = sync.RWMutex{} // read write mutex

func main() {
	go printHello() // becomes a go routine
	// when we make a go routine
	// the routine breaks off from the main routine
	// forms its own routine and executes there
	time.Sleep(100 * time.Millisecond) // to hold the main program
	// this holds the main program
	// so that printHello go routine can execute

	// using anonymous function
	go func() {
		fmt.Println("Hello Go for the second time")
	}()
	time.Sleep(100 * time.Millisecond) // to hold the main program

	// go routines can access variables in the outer scope
	msg := "Hello"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond) // to hold the main program

	// race condition in go routines
	// since go routines donot interrupt the normal execution of the main program
	msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond) // to hold the main program

	msg = "Hello"
	wg.Add(1)             // how many routines are added to waiting
	go func(msg string) { // by passing an argument
		fmt.Println(msg) // go will copy the value of the variable at the time its invoked
		wg.Done()        // specify that a routine ends so it decrements value from wg
	}(msg)
	msg = "Goodbye"
	wg.Wait() // wait for all routines to end

	runtime.GOMAXPROCS(100) // specify the number of OS thread

	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock() // make locking part of main thread
		go sayHello()
		mutex.Lock() // part of main thread
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello %v\n", counter)
	mutex.RUnlock() // unlock mutex after go routine ends
	wg.Done()
}

func increment() {
	counter++
	mutex.Unlock()
	wg.Done()
}

func printHello() {
	fmt.Println("Hello Go!")
}
