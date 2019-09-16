package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {
	switch num := number(); {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d less than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200\n", num)
	}
}
