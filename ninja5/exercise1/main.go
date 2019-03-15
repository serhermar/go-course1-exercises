package main

import "fmt"

type persona struct {
	nombre     string
	apellido   string
	saboresFav []string
}

func main() {
	p1 := persona{
		nombre:   "Gilito",
		apellido: "Donald",
		saboresFav: []string{
			"Fresa",
			"Merengada",
		},
	}

	p2 := persona{
		nombre:   "Pato",
		apellido: "Donald",
		saboresFav: []string{
			"Limon",
			"Chocolate",
		},
	}

	fmt.Println(p1.nombre, p1.apellido)
	for _, sabor1 := range p1.saboresFav {
		fmt.Printf("\t%v\n", sabor1)
	}

	fmt.Println(p2.nombre, p2.apellido)
	for _, sabor2 := range p2.saboresFav {
		fmt.Printf("\t%v\n", sabor2)
	}
}
