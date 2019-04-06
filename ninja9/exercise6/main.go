package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Operative System:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
}
