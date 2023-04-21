package main

import (
	"context"
	"log"
	"net"
	"test/grpc/cred/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)



var userInfo = map[string]int32 {
	"foo": 18,
	"bar": 20,
	"liql": 26,
	"hello": 100,
}

type Query struct {
	pb.UnimplementedQueryServer
}

func (q *Query) GetAge(ctx context.Context, info *pb.UserInfo) (*pb.AgeInfo, error) {
	age := userInfo[info.GetName()]
	var res = new(pb.AgeInfo)
	res.Age = age
	return res, nil
}

func (q *Query) GetName(ctx context.Context, info *pb.AgeInfo) (*pb.UserInfo, error) {
	age :=  info.GetAge()
	var res = new(pb.UserInfo)
	for key, value := range userInfo {
		if  age == value {
			res.Name = key
		}
	}
	return res, nil
}


const host = "0.0.0.0:1234"
func main() {
	// 创建socket监听器
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Panic(err)
	}

	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Panic(err)
	}

	// new一个gRPC服务器，用来注册服务
	grpcserver := grpc.NewServer(grpc.Creds(creds))
	// 注册服务方法
	pb.RegisterQueryServer(grpcserver, new(Query))
	// 开启gRPC服务
	err = grpcserver.Serve(listener)
	if err != nil {
		log.Panic(err)
	}
}
