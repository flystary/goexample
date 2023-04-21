package main

import (
	"fmt"
)

//单向通道

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

//  1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
//  2.<-chan int是一个只能接收的通道，可以接收但是不能发送。
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
