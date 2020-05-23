package main

import "fmt"

func main(){
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}

	// 通常，简短的声明也是定义信道的有效和简洁的方法
	// a := make(chan int)
	// fmt.Printf("Type of a is %T", a)
}