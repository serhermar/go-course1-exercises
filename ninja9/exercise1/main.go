package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func routine(count int) {
	fmt.Println("Hello from", count)
	wg.Done()
}

func main() {
	wg.Add(2)

	for i := 1; i <= 2; i++ {
		go routine(i)
	}

	wg.Wait()
}
