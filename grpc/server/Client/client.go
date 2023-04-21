package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"test/grpc/server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//建立无认证的连接
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	client := pb.NewQueryClient(conn)
	//返回GetAge方法对应的流用来接收来自服务端的message
	queryStream, _ := client.GetAge(context.Background(), &pb.UserInfo{Name: "foo"})
	// 开始接收message
	for {
		ageinfoRecv, err := queryStream.Recv()
		if err == io.EOF {
			log.Println("end of stream")
			break
		}
		fmt.Println(ageinfoRecv)
	}
}
