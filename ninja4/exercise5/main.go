package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(a[0:4], a[6:]...)
	fmt.Println(y)
}
