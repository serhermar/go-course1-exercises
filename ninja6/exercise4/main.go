package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) presentar() {
	fmt.Println("Hola soy", p.nombre, p.apellido, "y tengo", p.edad, "años")
}

func main() {
	p1 := persona{
		nombre:   "Sergio",
		apellido: "Hernández",
		edad:     35,
	}

	p1.presentar()
}
