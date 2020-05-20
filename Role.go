package main

import "fmt"

// Access rights program

const (
	isAdmin = iota << 1 // using bitwise shift
	canSeeAsia
	canSeeEurope
	canSeeAfrica
	canSeeNorthAmerica
	canSeeSouththAmerica
)

func main() {
	var role byte = isAdmin | canSeeAfrica

	fmt.Printf("User role %b\n", role)

	// to check if a role is valid using masking

	fmt.Printf("Is Admin? %v\n", role&isAdmin == isAdmin)
}
