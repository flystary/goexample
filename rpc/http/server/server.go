package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"net/rpc"
)


var userInfo = map[string]int {
	"foo": 12,
	"bar": 28,
	"hello": 100,
}

type Query struct {}

func (q *Query) GetAge(req string, res *string) error {
	*res = fmt.Sprintf("The age of %s is %d", req, userInfo[req])
	return nil
}

func (q *Query) GetName(req string, res *string) error {
	num, _ :=  strconv.Atoi(req)
	for key, value := range userInfo {
		if  num == value {
			*res = fmt.Sprintf("The age of %d is %s", req, key)
			return nil
		}
	}
	*res = fmt.Sprintf("The age of %s is Nil", req)
	return nil
}

func main() {
	if err := rpc.RegisterName("QueryService", new(Query)); err != nil {
		log.Println(err)
	}
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":1234", nil);err != nil {
		log.Panic(err)
	}
}
