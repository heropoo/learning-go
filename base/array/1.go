package main

import (
    "fmt"
)

func main(){
    var a [3]int
    fmt.Println(a)

    a[0] = 1
    a[1] = 3
    a[2] = 90
    fmt.Println(a)

    b := [3]int{12, 78, 50}
    fmt.Println(b)

    c := [3]int{12}
    fmt.Println(c)

    d := [...]int{12, 12, 231,44}
    fmt.Println(d)
}
