package main

import "fmt"

/**
*	Exercise 5: Custom Types and Type conversions
**/

// Section 1
type custom int

var x custom
var y int

func main() {
	// Section 2
	fmt.Println(x)        // a
	fmt.Printf("%T\n", x) // b
	x = 42                // c
	fmt.Println(x)        // d

	// Section 3
	y = int(x)
	fmt.Println(y)
}
