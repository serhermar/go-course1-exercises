package main

import "fmt"

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "Hello"
}

func main() {
	fmt.Println("foo Function: ", foo())

	x, y := bar()

	fmt.Println("bar Function: ", x, y)
}
