package main

import (
	"fmt"
	"math"
)

type cuadrado struct {
	longitud float64
}

func (r cuadrado) area() float64 {
	return r.longitud * r.longitud
}

type circulo struct {
	radio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

type figura interface {
	area() float64
}

func info(f figura) {
	fmt.Println("El area del es", f.area())
}

func main() {
	circle := circulo{
		radio: 10,
	}

	square := cuadrado{
		longitud: 10,
	}

	info(circle)
	info(square)
}
