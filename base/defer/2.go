package main

import "fmt"

func main() {
	name := "Seven"
	fmt.Println("Original String: %s\n", string(name))
	fmt.Println("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
