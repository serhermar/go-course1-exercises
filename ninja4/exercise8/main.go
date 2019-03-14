package main

import "fmt"

func main() {
	x := map[string][]string{
		`adam`:    []string{`play`, `soccer`},
		`jhon`:    []string{`relax`, `karate`},
		`jameson`: []string{`play`, `videogames`},
	}

	for k, v := range x {
		fmt.Println("Record:", k)

		for i, v2 := range v {
			fmt.Printf("\tColumn %d - %v\n", i, v2)
		}
	}
}
