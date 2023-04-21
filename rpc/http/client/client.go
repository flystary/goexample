package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	age  := flag.String("a", "0", "age")
	name := flag.String("n", "","name")

	flag.Parse()
	// fmt.Println(age, name)

	client, _ := rpc.DialHTTP("tcp", ":1234")
	var res string
	if *age != "" {
		err := client.Call("QueryService.GetName", *age, &res)
		if err != nil {
			log.Panicln(err)
		}
	}
	if *name != "" {
		err := client.Call("QueryService.GetAge", *name, &res)
		if err != nil {
			log.Panicln(err)
		}
	}

	fmt.Println(res)
	_ = client.Close()
}
