package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//fmt.Println("Hello World")

	http.HandleFunc("/", showIndex)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and serve: ", err.Error())
	}
	fmt.Println("Listen and serve on 127.0.0.1:8080")
}

//首页
func showIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

// 登录
func login(w http.ResponseWriter, r *http.Request) {

}
