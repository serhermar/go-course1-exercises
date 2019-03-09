package main

import "fmt"

/**
*	Exercise 3: Assignement Operator and Sprintf
**/

// Section 1
var x int = 42              // a
var y string = "James Bond" // b
var z bool = true           // c

func main() {
	// Section 2
	s := fmt.Sprintf("%v\t%s\t%t", x, y, z) // a
	fmt.Println(s)                          // b
}
