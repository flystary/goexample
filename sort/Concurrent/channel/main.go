package main

import (
	"fmt"
)

// var 变量 chan 元素类型
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	//	ch := make(chan int)
	//	go recv(ch)
	//	ch <- 10
	//	fmt.Println("发送成功")
	//	chans()
	chans1()
}

func chans() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}

func chans1() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	for i := range ch2 {

		fmt.Printf("%v\n", i)
	}
}
