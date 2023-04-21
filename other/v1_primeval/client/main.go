package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err.Error())
	}
	req := "mclik"
	var resp *string
	err = client.Call("EncryptionUtil.Encryption", req, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)

	syncCall := client.Go("EncryptionUtil.Encryption", req, &resp, nil)
	replayDone := <-syncCall.Done
	fmt.Println(replayDone)
	fmt.Println(*resp)
}
