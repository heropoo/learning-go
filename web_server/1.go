package main

import (
    "io"
    "log"
    "net/http"
    "fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello, world!")
}
func main() {
    http.HandleFunc("/hello", helloHandler)
    fmt.Println("Http server is listen at 127.0.0.1:8080")
    err := http.ListenAndServe("http://127.0.0.1:8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
    
}
