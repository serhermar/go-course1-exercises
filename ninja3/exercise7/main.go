package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("Mod(%d / 4) = 0\n", i)
		} else if i%4 == 1 {
			fmt.Printf("Mod(%d / 4) = 1\n", i)
		} else if i%4 == 2 {
			fmt.Printf("Mod(%d / 4) = 2\n", i)
		} else {
			fmt.Printf("Mod(%d / 4) = 3\n", i)
		}
	}
}
