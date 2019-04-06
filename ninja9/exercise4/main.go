package main

import (
	"fmt"
	"sync"
)

var increment int
var wg sync.WaitGroup
var m sync.Mutex

/*
* With this function we create race conditions
 */
func routine() {
	m.Lock()
	read := increment
	read++
	increment = read
	m.Unlock()
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
