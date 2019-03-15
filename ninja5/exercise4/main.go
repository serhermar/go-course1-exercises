package main

import "fmt"

type vehiculo struct {
	puertas int
	color   string
}

func main() {
	s := struct {
		puertas int
		color   string
	}{
		puertas: 5,
		color:   "Gris",
	}

	fmt.Println("Puertas: ", s.puertas)
	fmt.Println("Color: ", s.color)
}
