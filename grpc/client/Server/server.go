package main

import (
	"io"
	"log"
	"net"
	"test/grpc/client/pb"

	"google.golang.org/grpc"
)

const host = ":1234"

var userInfo = map[string]int32{
	"foo": 19,
	"bar": 29,
	"lql": 30,
}

type Query struct {
	pb.UnimplementedQueryServer
}

func (q *Query) GetAge(serverStream pb.Query_GetAgeServer) error {
	log.Println("start")
	var names_received []*pb.UserInfo
	for {
		userinfoRecv, err := serverStream.Recv()
		if err == io.EOF {
			log.Println("end")
			break
		}
		log.Printf("Name: %s\n", userinfoRecv.GetName())
		names_received = append(names_received, userinfoRecv)
	}

	var ages_sum int32
	for _, v := range names_received {
		ages_sum += userInfo[v.GetName()]
	}

	log.Printf("send message about the total of ages:%d\n", ages_sum)
	err := serverStream.SendAndClose(&pb.AgeInfo{Age: ages_sum})
	if err != nil {
		log.Panic(err)
	}
	log.Println("end of the send direction of stream")
	return nil
}

func main() {
	listener, _ := net.Listen("tcp", host)
	grpcserver := grpc.NewServer()
	pb.RegisterQueryServer(grpcserver, new(Query))

	_ = grpcserver.Serve(listener)
}
