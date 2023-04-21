package main

import (
	"context"
	"log"
	"net"
	"test/grpc/base/pb"

	"google.golang.org/grpc"
)


var userInfo = map[string]int32{
	"foo": 18,
	"bar": 20,
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

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panic(err)
	}

	grpcserver := grpc.NewServer()
	pb.RegisterQueryServer(grpcserver, new(Query))
	err = grpcserver.Serve(listener)
	if err != nil {
		log.Panic(err)
	}
}