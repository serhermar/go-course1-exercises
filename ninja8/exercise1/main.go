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
	u1 := usuario{
		Nombre: "Sergio",
		Edad:   36,
	}

	u2 := usuario{
		Nombre: "Fermina",
		Edad:   35,
	}

	usuarios := []usuario{u1, u2}

	json, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json))

}
