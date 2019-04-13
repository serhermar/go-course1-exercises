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

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Error building JSON: %v", err)
	}

	return bs, nil
}
