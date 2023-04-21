package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"tcp/init-2/proto"
)

func main() {
	// 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("服务器启动失败....", err)
		return
	}
	fmt.Println("监听成功...")
	// 等待连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接建立失败...", err)
			break
		}
		fmt.Println("连接成功...")

		go Process(conn)
	}
	// 通信
}

func Process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := proto.Decode(reader)
		fmt.Println("收到消息：", msg)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode失败，err:", err)
		}
	}
}
