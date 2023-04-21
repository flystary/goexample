package main

import (
	"io"
	"log"
	"net"
	"test/grpc/bidirectional/pb"

	"google.golang.org/grpc"
)

var userInfo = map[string]int32 {
	"liql": 18,
	"bar": 20,
	"foo": 10,
}

type Query struct {
	pb.UnimplementedQueryServer
}



func (q *Query) GetAge(ServerStream pb.Query_GetAgeServer) error {
	log.Println("---> start")
	for {
		userinfoRecv, err := ServerStream.Recv()
		if err == io.EOF {
			log.Println("end--->")
			break
		}
		// log.Printf("The name of user received is %s", userinfoRecv.GetName())
		// 返回响应message
		log.Printf(">---%s---<", userinfoRecv.GetName())
		err = ServerStream.Send(&pb.AgeInfo{Age: userInfo[userinfoRecv.Name]})
		if err != nil {
			log.Panic(err)
		}
	}

	return nil
}


func main() {
	listener, _ := net.Listen("tcp", ":1234")
	server := grpc.NewServer()
	pb.RegisterQueryServer(server, new(Query))
	_ =server.Serve(listener)
}
