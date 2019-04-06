package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var increment int64
var wg sync.WaitGroup

func routine() {
	atomic.AddInt64(&increment, 1)
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
