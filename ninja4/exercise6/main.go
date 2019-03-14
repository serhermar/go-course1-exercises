package main

import "fmt"

func main() {
	states := make([]string, 50, 50)

	states = []string{
		"Alabama",
		"Alaska",
		"Arizana",
	}

	fmt.Println("Len: ", len(states))
	fmt.Println("Cap: ", cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}
}
