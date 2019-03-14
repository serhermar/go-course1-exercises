package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range a {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	fmt.Printf("%T\n", a)
}
