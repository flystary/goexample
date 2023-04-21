package main

import (
	"test/PRC/P/spider"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"net"
	"net/http"
)

type server struct{}

const (
	// Address 监听地址
	Address string = "localhost:8080"
	// Method 通信方法
	Method string = "tcp"
)

// GetAddressResponse 接收client端的请求,函数名需保持一致
// ctx参数必传
// 参数二为自定义的参数,需从pb文件导入,因此pb文件必须可导入,文件放哪里随意
// 返回值同参数二,为pb文件的返回结构体指针
func (s *server) GetAddressResponse(ctx context.Context, a *spider.SendAddress) (*spider.GetResponse, error) {
	// 逻辑写在这里
	switch a.Method {
	case "get", "Get", "GET":
		// 演示微服务用,故只写get示例
		status, body, err := get(a.Address)
		if err != nil {
			return nil, err
		}
		res := spider.GetResponse{
			HttpCode: int32(status),
			Response: body,
		}
		return &res, nil
	}
	return nil, nil
}

func get(address string) (s int, r string, err error) {
	// get请求
	resp, err := http.Get(address)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	s = resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	r = string(body)
	return
}

func main() {
	// 监听本地端口
	listener, err := net.Listen(Method, Address)
	if err != nil {
		return
	}
	s := grpc.NewServer()                       // 创建GRPC
	spider.RegisterGoSpiderServer(s, &server{}) // 在GRPC服务端注册服务

	reflection.Register(s) // 在GRPC服务器注册服务器反射服务
	// Serve方法接收监听的端口,每到一个连接创建一个ServerTransport和server的grroutine
	// 这个goroutine读取GRPC请求,调用已注册的处理程序进行响应
	err = s.Serve(listener)
	if err != nil {
		return
	}
}
