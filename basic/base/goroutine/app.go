package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a NewTask")
		time.Sleep(time.Second)
	}
}

//world
//world
//hello
//hello
func str() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func strs() {
	now := time.Now().Format("15:04:05")
	go func() {
		defer fmt.Println("A.defer", now)
		func() {
			defer fmt.Println("B.defer", now)
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer", now)
			fmt.Println("B", now)
		}()
		fmt.Println("A", now)
	}()
	for {
	}
}
