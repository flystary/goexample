package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接失败,err:", err)
		return
	}
	defer conn.Close()
	msg := "hello socket"
	for i := 0; i < 20; i++ {
		conn.Write([]byte(msg))
	}
}

