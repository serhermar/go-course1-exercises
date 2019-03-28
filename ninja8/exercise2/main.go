package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func main() {
	j := `[{"nombre":"Sergio","edad":36},{"nombre":"Fermina","edad":35}]`

	usuarios := []usuario{}

	err := json.Unmarshal([]byte(j), &usuarios)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Edad)
	}
}
