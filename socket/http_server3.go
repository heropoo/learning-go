package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
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
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	fmt.Print("-")
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // 设置2分钟超时
	defer conn.Close()                                    // 退出前关闭连接
	request := make([]byte, 128)
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 {
			break
		} else if strings.TrimSpace(string(request[:10])) == "GET /test" {
			msg := "HTTP/1.1 200 OK\r\n"
			msg += "Content-Length: 10\r\n"
			msg += "\r\n"
			msg += "Hello Test"

			conn.Write([]byte(msg))
		} else {
			msg := "HTTP/1.1 200 OK\r\n"
			msg += "Content-Length: 11\r\n"
			msg += "\r\n"
			msg += "Hello World"

			conn.Write([]byte(msg))
		}
		request = make([]byte, 128) //清除上次读取的内容
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
