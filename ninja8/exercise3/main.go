package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	err := json.NewEncoder(os.Stdout).Encode(usuarios)
	if err != nil {
		fmt.Println(err)
	}
}
