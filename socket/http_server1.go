package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("get a connect")
		msg := "HTTP/1.1 200 OK\r\n"
		msg += "Content-Length: 12\r\n"
		msg += "\r\n"
		msg += "Hello World\n"
		conn.Write([]byte(msg))
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
