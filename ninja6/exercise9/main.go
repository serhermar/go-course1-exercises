package main

import "fmt"

func foo(f func(xi []int) int, si []int) int {
	n := f(si)
	n++
	return n
}

func suma(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(foo(suma, si))
}
