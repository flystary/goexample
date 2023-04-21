package main

import (
	"fmt"
	"time"
)

//全局变量，创建一个channel
var ch = make(chan int)

// Printer 打印机
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

// Person1 person1执行完成之后，person2执行
func Person1() {
	Printer("hello")
	ch <- 666
}

func Person2() {
	<-ch
	Printer("world")
}

func main() {
	go Person1()
	go Person2()
	for {
	}

}
