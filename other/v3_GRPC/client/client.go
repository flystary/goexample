package main

import (
	"test/PRC/v3_GRPC/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	encryptionServiceClient := message.NewEncryptionServiceClient(conn)
	orderRequest := &message.EncryptionRequest{Str: "mclik"}
	result, err := encryptionServiceClient.Encryption(context.Background(), orderRequest)
	if result != nil {
		fmt.Println(result.Result)
	}
}
