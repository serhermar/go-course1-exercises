package main

import "fmt"

type customError struct {
	info string
}

func (ce customError) Error() string {
	return fmt.Sprintf("Error: %v", ce.info)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	e1 := customError{
		info: "Error1",
	}

	foo(e1)
}
