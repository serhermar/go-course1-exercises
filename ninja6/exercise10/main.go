package main

import "fmt"

func foo() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func main() {
	f := foo()

	a := f()
	b := f()
	c := f()
	d := f()

	fmt.Println(a, b, c, d)

}
