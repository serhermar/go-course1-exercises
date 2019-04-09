package main

import "fmt"

/* Solution with anonymous function in gorutine
//
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
*/

/* Solution with buffered channel */
func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
