package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"test/grpc/base/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	age  := flag.String("a", "", "age")
	name := flag.String("n", "","name")
	flag.Parse()

	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()
	client := pb.NewQueryClient(conn)
	if *name != "" {
		age, _ := client.GetAge(context.Background(), &pb.UserInfo{Name: *name})
		fmt.Println(age)
	}

	if *age != "" {
		I32 :=  stringInt32(*age)
		name, _ := client.GetName(context.Background(), &pb.AgeInfo{Age: I32})
		fmt.Println(name)
	}
}

func stringInt32(s string) int32 {
	Int64,err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Panic(err)
	}
    Int32 := int32(Int64)
	return Int32
}