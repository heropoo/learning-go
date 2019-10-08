package main

import "fmt"

func main(){
    a := [5]int{132, 535, 64, 6, 90}
    var b []int = a[1:4]
    fmt.Println(b)
}
