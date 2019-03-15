package main

import "fmt"

type vehiculo struct {
	puertas int
	color   string
}

type sedan struct {
	vehiculo
	lujoso bool
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

func main() {
	s := sedan{
		vehiculo: vehiculo{
			puertas: 5,
			color:   "verde",
		},
		lujoso: true,
	}

	c := camion{
		vehiculo: vehiculo{
			puertas: 5,
			color:   "verde",
		},
		cuatroRuedas: false,
	}

	fmt.Println("------ Sedan ------")
	fmt.Println("Puertas: ", s.puertas)
	fmt.Println("Color: ", s.color)
	fmt.Println("Lujoso?: ", s.lujoso)

	fmt.Println("------ Camion ------")
	fmt.Println("Puertas: ", c.puertas)
	fmt.Println("Color: ", c.color)
	fmt.Println("Cuatro Ruedas?: ", c.cuatroRuedas)
}
