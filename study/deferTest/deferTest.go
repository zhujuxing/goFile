package main

import (
	"fmt"
	"net"
)

func topSend() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err == nil {
		defer conn.Close()
		fmt.Println("remote address:", conn.RemoteAddr())
	} else {
		fmt.Println("error", err)
	}
}

func main() {
	topSend()
}
