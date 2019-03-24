package main

import "fmt"

func foo(s ...int) int {
	total := 0

	for _, v := range s {
		total += v
	}

	return total
}

func bar(s []int) int {
	total := 0

	for _, v := range s {
		total += v
	}

	return total
}

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum1 := foo(si...)
	sum2 := bar(si)

	fmt.Println(sum1, sum2)
}
