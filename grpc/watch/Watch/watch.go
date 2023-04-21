package main

import (
	"context"
	"io"
	"log"
	"test/grpc/watch/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	client := pb.NewQueryClient(conn)
	stream, _ := client.Watch(context.Background(), &pb.WatchTime{Time: 1000})
	for {
		userInfoRecv, err := stream.Recv()
		if err == io.EOF {
			log.Panicln("end of watch")
			break
		} else if err != nil {
			log.Panicln(err)
			break
		}
		log.Printf("The Name of %s is UPDATED\n", userInfoRecv.GetName())
	}
}