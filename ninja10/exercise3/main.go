package main

import "fmt"

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func recibir(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	c := gen()

	recibir(c)

	fmt.Println("Finishing")
}
