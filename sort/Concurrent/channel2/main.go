package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	dataChan := make(chan int, 100)
	go func() {
		for {
			select {
			case data := <-dataChan:
				fmt.Println("data", data)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	for i := 0; i < 100; i++ {
		dataChan <- i
	}

	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
