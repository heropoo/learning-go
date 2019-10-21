package main

import "fmt"

func mutate(s []rune) string {
	s[0] = '👋'
	return string(s)
}

func main() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
