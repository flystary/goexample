package main

import (
	"context"
	"fmt"
	"log"
	"test/grpc/client/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const host = ":1234"

func main() {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()
	client := pb.NewQueryClient(conn)
	queryStream, _ := client.GetAge(context.Background())
	_ = queryStream.Send(&pb.UserInfo{Name: "foo"})
	time.Sleep(time.Second)
	_ = queryStream.Send(&pb.UserInfo{Name: "bar"})
	time.Sleep(time.Second)

	ages_sum, err := queryStream.CloseAndRecv()

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("The total of ages of (foo + bar) is %d\n", ages_sum.GetAge())
}
