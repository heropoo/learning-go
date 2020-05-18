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
	http.HandleFunc("/logout", logout)

	fmt.Println("Listen and serve on http://127.0.0.1:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Listen and serve: ", err.Error())
	}
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

//注销
func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout")
}
