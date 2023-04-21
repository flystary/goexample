package main

import (
	"fmt"
	"sync"
)

//互斥锁
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //解锁
	}
	wg.Done()
}

func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
