package main

import (
	"context"
	"io"
	"log"
	"test/grpc/bidirectional/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const host = ":1234"

func main() {
	conn, _ := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer conn.Close()
	client := pb.NewQueryClient(conn)

	log.Println("start")
	queryStream, _ := client.GetAge(context.Background())

	ch := make(chan string, 2)

	go func() {
		names := []string{"foo", "bar"}
		for _, v := range names {
			log.Printf("send Name: %s\n", v)
			ch <- v
			_ = queryStream.Send(&pb.UserInfo{Name: v})
			time.Sleep(time.Second)
		}

		err := queryStream.CloseSend()
		if err != nil {
			log.Println(err)
		}
		close(ch)
	}()

	for {
		name := <-ch
		ageinfoRecv, err := queryStream.Recv()
		if err == io.EOF {
			log.Println("end of Stream")
			break
		}

		log.Printf("The age of %s is %d\n", name, ageinfoRecv.GetAge())
	}
}
