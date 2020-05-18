package main

import "fmt"

func main() {
	n := 20
	nx := n
	for i := 0; i < n; i++ {
		if nx > 0 {
			for nxi := 0; nxi < nx; nxi++ {
				fmt.Print(" ")
			}
		}

		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
		nx--
	}
}
