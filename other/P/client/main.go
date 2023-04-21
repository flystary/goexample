package main

import (
	"test/PRC/P/spider"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

const (
	// Address server端地址
	Address string = "localhost:8080"
)

func main() {

	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 连接GRPC
	c := spider.NewGoSpiderClient(conn)

	// 创建要发送的结构体
	req := spider.SendAddress{
		Address: "http://www.baidu.com",
		Method:  "get",
	}

	// 调用server的注册方法
	r, err := c.GetAddressResponse(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印返回值
	fmt.Println(r)
}
