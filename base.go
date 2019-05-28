// base
package main

import (
	"fmt"
)

func run() {
	fmt.Print("hello world")
}

func main() {
	go run()
}
