package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello World Goroutine")
}

func main() {
	go hello()
	fmt.Println("Main function")
}
