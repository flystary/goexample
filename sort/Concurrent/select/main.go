package main

import (
	"fmt"
	"time"
)

//select可以同时监听一个或多个channel，直到其中一个channel ready

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 4)
	ch <- "test2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go test1(output1)
	go test2(output2)

	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}
