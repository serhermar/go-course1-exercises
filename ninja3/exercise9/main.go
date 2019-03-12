package main

import "fmt"

func main() {

	deporteFav := "Karate"

	switch deporteFav {
	case "Futbol":
		fmt.Println("Soccer!")
	case "Karate":
		fmt.Println("Give me 5!")
	default:
		fmt.Println("Do you like books?")
	}

}
