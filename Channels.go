package main

import (
	"fmt"
	"sync"
)

// channels are tools which can be used to pass messages between 2 routines

var wg = sync.WaitGroup{}

func main() {
	// creating a channel
	c := make(chan int) // channels are strongly typed

	wg.Add(2)
	// go routine to recieve a data
	go func() {
		i := <-c
		fmt.Println(i)
		wg.Done()
	}()
	// channel to send data
	go func() {
		c <- 42
		wg.Done()
	}()
	wg.Wait()

	unbuffered channels stop the execution of the go routine

	ch := make(chan int) // unbuff so only 1 data can exist is the channel

	go func() { // go routine to process data
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	for j := 0; j < 5; j++ { // send data to channel 2 times
		wg.Add(2)
		go func() {
			ch <- 42 // send data to channel, execution stops until data is taken from channel
			wg.Done()
		}()
	}
	// there is no routine to process the data pushed into the channel
	// so the program goes to a deadlock as the channel needs to be freed
	wg.Wait()

	// having 2 go routines communicate as readers and writers
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch // as reader
		fmt.Println(i)
		ch <- 27 // as writer
		wg.Done()
	}()
	go func() {
		ch <- 42          // as writer
		fmt.Println(<-ch) // as reader
		wg.Done()
	}()
	wg.Wait()

	// specifying only 1 way channels
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) { // can only recieve data from channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // can only send data to channel
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()

	// buffered channels and looping over channels
	ch := make(chan int, 50) // create a channel with buffer storage 50
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch { // looping over the channel
			fmt.Println(i)
		} // loop knows that the channel has stoped exec
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42  // channel can now
		ch <- 27  // recieve multiple data
		ch <- 17  // because of buffer channel
		close(ch) // closing the channel
		wg.Done()
	}(ch)
	wg.Wait()

	// using the ,ok syntax to process loops
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				fmt.Println("Channel closed")
				// this works best when
				// rather than breaking
				// from the loop routine you
				// have to keep listening
				// for incomming channel data
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		ch <- 17
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
