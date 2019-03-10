package main

import "fmt"

const (
	actualYear = 2019
	a          = actualYear + iota
	b          = actualYear + iota
	c          = actualYear + iota
	d          = actualYear + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
