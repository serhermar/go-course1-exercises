package main

import "fmt"

type persona struct {
	nombre string
}

func cambiame(p *persona, nombre string) {
	p.nombre = nombre

	// (*p).nombre = nombre
}

func main() {
	p := persona{
		nombre: "sergio",
	}

	fmt.Println("Antes: ", p)

	cambiame(&p, "juan")

	fmt.Println("Despues: ", p)
}
