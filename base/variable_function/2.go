package main

import (
	"fmt"
)

func find(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}

	if !found {
		fmt.Println(num, "not found in ", nums)
	}

	fmt.Printf("\n")
}

func main() {
	find(1, []int{2, 3, 1, 45325, 3})
	find(1, []int{2, 3, 45325, 3})
	find(1, []int{})
}
