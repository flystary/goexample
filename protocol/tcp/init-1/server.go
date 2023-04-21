package main

import (
	"fmt"
	"net"
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
	var temp [128]byte
	for {
		n, err := conn.Read(temp[:])
		if err != nil {
			fmt.Println("服务端读取数据失败....")
			return
		}
		fmt.Println("读取数据成功:")
		fmt.Println(string(temp[:n]))
	}
}
