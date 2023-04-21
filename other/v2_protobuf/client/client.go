package main

import (
	"test/PRC/v2_protobuf/message"
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}
	req := message.EncryptionRequest{Str: "mclik"}
	resp := message.EncryptionResult{}
	//同步调用
	err = client.Call("EncryptionService.Encryption", &req, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp.GetResult())
}
