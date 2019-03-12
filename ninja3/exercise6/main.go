package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("Mod(%d / 4) = %d\n", i, i%4)
		}
	}
}
