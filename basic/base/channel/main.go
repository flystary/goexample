package main

import (
	"fmt"
	"time"
)

//打印机
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

//person1执行完成之后，person2执行
func person1() {
	Printer("hello")
}
func person2() {
	Printer("world")
}

func main() {
	go person1()
	go person2()
	for {
	}
}
