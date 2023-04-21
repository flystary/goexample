package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"test/grpc/watch/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func stringInt32(s string) int32 {
	Int64,err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Panic(err)
	}
    Int32 := int32(Int64)
	return Int32
}

func main() {
	//建立无认证的连接
	age  := flag.String("a", "", "age")
	name := flag.String("n", "","name")
	flag.Parse()

	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()
	client := pb.NewQueryClient(conn)

	//RPC方法调用
	ctx := context.Background()

	if *name != "" {
		age, _ := client.GetAge(context.Background(), &pb.UserInfo{Name: *name})
		log.Printf("Before updating, the age is %d\n", age.GetAge())
		//更新年龄
		log.Println("updating")
		client.Update(ctx, &pb.UserInfo{Name: *name})
		//再获取更新后的年龄
		age, _ = client.GetAge(ctx, &pb.UserInfo{Name: *name})
		log.Printf("After updating, the age is %d\n", age.GetAge())
	}

	if *age != "" {
		I32 :=  stringInt32(*age)
		name, _ := client.GetName(context.Background(), &pb.AgeInfo{Age: I32})
		fmt.Println(name)
	}
}
