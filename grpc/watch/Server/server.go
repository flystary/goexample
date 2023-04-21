package main

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"
	"test/grpc/watch/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)


var userInfo = map[string]int32 {
	"foo": 12,
	"bar": 20,
	"liql": 26,
	"hello": 100,
}

type Query struct {
	mu		sync.Mutex
	ch 		chan string
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

func (q *Query) Update(ctx context.Context, info *pb.UserInfo) (*emptypb.Empty, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	name := info.GetName()
	userInfo[name] += 1
	if q.ch != nil {
		q.ch <- name
	}
	return &emptypb.Empty{}, nil
}

func (q *Query) Watch(timeSpcify *pb.WatchTime, ServerStream pb.Query_WatchServer) error {
	if q.ch != nil {
		return errors.New("Watching is running, please stop first")
	}
	q.ch = make(chan string, 1)

	for {
		select {
		case <-time.After(time.Second * time.Duration(timeSpcify.GetTime())):
			close(q.ch)
			q.ch = nil
			return nil
	case nameModify := <-q.ch:
			log.Printf("The Name of %s is UPDATE\n", nameModify)
			ServerStream.Send(&pb.UserInfo{Name: nameModify})
		}
	}
}

func main() {
	// 创建socket监听器
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panic(err)
	}
	// new一个gRPC服务器，用来注册服务
	grpcserver := grpc.NewServer()
	// 注册服务方法
	pb.RegisterQueryServer(grpcserver, &Query{})
	// 开启gRPC服务
	err = grpcserver.Serve(listener)
	if err != nil {
		log.Panic(err)
	}
}