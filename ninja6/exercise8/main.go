package main

import "fmt"

func zeroToNine() func() {
	zeroToNineFunc := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	return zeroToNineFunc
}

func main() {
	a := zeroToNine()

	a()
}
