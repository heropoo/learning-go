package main

import "fmt"

func main() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}

	personSalary["mike"] = 9000
	fmt.Println("map before deletion", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)
}
