package main

import "fmt"

func main() {
	actualYear := 2019
	bornYear := 1982

	year := bornYear

	for {
		if year >= actualYear {
			break
		}

		fmt.Println(year)

		year++
	}
}
