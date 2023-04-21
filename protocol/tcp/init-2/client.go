package main

import (
	"fmt"
	"net"
	"tcp/init-2/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接失败,err:", err)
		return
	}
	defer conn.Close()
	msg := "hello socket"

	for i := 0; i < 4; i++ {
		b, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("Encode失败, err: ", err)
			continue
		}
		conn.Write(b)
		fmt.Println("发送成功...,msg：", b)
	}
}
