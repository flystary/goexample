package main

import (
	"fmt"
	"time"
)

//// 判断管道有没有存满
func main() {
	output1 := make(chan string, 10)
	go write(output1)

	for s := range output1 {
		fmt.Println("res", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello woeo":
			fmt.Println("write hello")
		default:
			fmt.Println("chan full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
