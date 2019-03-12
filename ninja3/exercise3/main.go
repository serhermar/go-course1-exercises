package main

import "fmt"

func main() {
	actualYear := 2019
	bornYear := 1982

	year := bornYear

	for year < actualYear {
		fmt.Println(year)

		year++
	}
}
