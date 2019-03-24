package main

import "fmt"

func foo() {
	fmt.Println("Hello from foo")
}

func bar() {
	fmt.Println("Hello from bar")
}

func main() {
	defer foo()
	bar()
}
