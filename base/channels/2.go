package main

import "fmt"

func hello(done chan bool){
	fmt.Println("Hello world goroutine")
	done <- true
}

func main(){
	//var res chan bool
	done := make(chan bool)
	go hello(done)
	<- done
	//fmt.Printf("type of res is %T", res)
	fmt.Println("main function")
}