package main

import (
	"time"
	"fmt"
)

func hello(done chan bool){
	fmt.Println("Hello go goroutine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("Hello go goroutine awake and going to write to done")
	done <- true
}

func main(){
	//var res chan bool
	done := make(chan bool)
	fmt.Println("Main going to call helle go goroutine")
	go hello(done)
	<- done
	fmt.Println("main received data")
}