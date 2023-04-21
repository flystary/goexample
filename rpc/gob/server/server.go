package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)


var userInfo = map[string]int {
	"foo": 12,
	"bar": 28,
}

type Query struct {}

func (q *Query) GetAge(req string, res *string) error {
	*res = fmt.Sprintf("The age of %s is %d", req, userInfo[req])
	return nil
}

func main() {
	// 注册服务方法
	if err := rpc.RegisterName("QueryService", new(Query)); err != nil {
		log.Println(err)
	}

	listener, _ := net.Listen("tcp", ":1234")
	rpc.Accept(listener)

}