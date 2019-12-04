package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
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
	data, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	io.WriteString(w, string(data))
}

// 登录
func login(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Login")
}
