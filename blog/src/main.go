package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonData struct {
	Code int
	Msg  string
	data string
}

func main() {
	http.HandleFunc("/", showIndex)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	fmt.Println("Listen and serve on http://127.0.0.1:8080")

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	//err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Listen and serve: ", err.Error())
	}
}

//首页
func showIndex(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("../public/index.html")
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
	//fmt.Fprintf(w, "Logout")
	d := &JsonData{}
	d.Code = 0
	d.Msg = "Ok"
	d.data = "logout"
	data, _ := json.Marshal(d)

	w.Header().Add("Content-Type", "application/json")
	//fmt.Fprintf(w, string(data))
	w.Write([]byte(data))
}
