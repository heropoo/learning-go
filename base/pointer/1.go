package main

import "fmt"

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Printf("Address of b is %p\n", a)
	fmt.Println(&a)
}
