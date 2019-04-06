package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

/*
* 	(t T) 	---> T and *T
*	(t *T) 	---> *T
 */
func main() {
	p := person{
		name: "Sergio",
	}

	saySomething(&p)
	// saySomething(p) // Cannot use p instead of *p
}
