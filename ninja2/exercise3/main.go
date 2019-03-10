package main

import "fmt"

const (
	pi         = 3.1416
	msg string = "Hello World!"
)

func main() {
	fmt.Printf("Type %T Value %f\n", pi, pi)
	fmt.Printf("Type %T Value %s\n", msg, msg)
}
