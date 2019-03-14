package main

import "fmt"

func main() {
	batman := []string{"Batman", "Boss", "Black Dress"}
	robin := []string{"Robin", "Helper", "Color Dress"}

	x := [][]string{batman, robin}

	for i, v := range x {
		fmt.Println("Record ", i)

		for j, v2 := range v {
			fmt.Printf("\tColumn %d: %v\n", j, v2)
		}
	}
}
