package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name    string
	Surname string
	Phrases []string
}

func main() {
	p1 := person{
		Name:    "Sergio",
		Surname: "Surname",
		Phrases: []string{"Hello", "Bye"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
