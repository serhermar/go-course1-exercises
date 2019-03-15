package main

import "fmt"

type persona struct {
	nombre     string
	apellido   string
	saboresFav []string
}

func main() {
	p1 := persona{
		nombre:   "Tito",
		apellido: "Gilito",
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

	personas := map[string]persona{
		p1.apellido: p1,
		p2.apellido: p2,
	}

	for _, p := range personas {
		fmt.Println(p.nombre, p.apellido)
		for _, sabor := range p.saboresFav {
			fmt.Printf("\t%v\n", sabor)
		}
	}
}
