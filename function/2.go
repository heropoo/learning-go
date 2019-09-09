package main

import(
    "fmt"
)

func main() {
    area, _ := rectProps(10.8, 5.6)
    fmt.Printf("Area %f", area)
}

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}
