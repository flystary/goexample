package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", ":1234")

	var res string
	err := client.Call("QueryService.GetAge", "foo", &res)
	if err != nil {
		log.Println(err)
	}
	
	fmt.Println(res)

	_ = client.Close()
}