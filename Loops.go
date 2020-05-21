package main

import "fmt"

func main() {
	// looping

	// simple for loop
	for i := 0; i < 5; i += 2 {
		fmt.Println(i)
	}

	// double var
	for i, j := 0, 0; i < 5; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)

	// infinite loop
	for {
		fmt.Println(i)
		if i > 10 {
			break
		}
		i++
	}

	// working with collections

	s := []string{"Salute", "To", "The", "World"}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
