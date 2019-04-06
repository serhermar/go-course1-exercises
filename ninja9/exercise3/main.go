package main

import (
	"fmt"
	"runtime"
	"sync"
)

var increment int
var wg sync.WaitGroup

/*
* With this function we create race conditions
 */
func routine() {
	read := increment

	runtime.Gosched()
	read++

	increment = read
	wg.Done()
}

func main() {

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go routine()
	}

	wg.Wait()

	fmt.Println(increment)
}
