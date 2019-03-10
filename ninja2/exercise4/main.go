package main

import "fmt"

func main() {
	a := 50
	fmt.Printf("%d %b %#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d %b %#x\n", b, b, b)
}
