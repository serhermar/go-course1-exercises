package main

import "fmt"

/**
*	Exercise 4: Custom Types
**/

// Section 1
type custom int

// Section 2
var x custom

func main() {
	// Section 3
	fmt.Println(x)        // a
	fmt.Printf("%T\n", x) // b
	x = 42                // c
	fmt.Println(x)        // d
}
