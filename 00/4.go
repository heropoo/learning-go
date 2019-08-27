package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 123.9, 343.1
	c := math.Min(a, b)
	fmt.Println("minimum value is", c)
}
