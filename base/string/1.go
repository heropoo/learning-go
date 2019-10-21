package main

import "fmt"

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:%c ", s[i], s[i])
	}
}

func main() {
	name := "Hello world"
	printBytes(name)
}
