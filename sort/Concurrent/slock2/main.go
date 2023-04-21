package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁分为两种：读锁和写锁。
//当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
//当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	//lock.Lock()  //加互斥锁
	rwlock.Lock() //加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) //假设读操作耗时10毫秒
	rwlock.Unlock()                   //解写锁
	//lock.Unlock()					//解互斥锁
	wg.Done()
}

func read() {
	//lock.Lock()
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	//lock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()

	fmt.Println(end.Sub(start))
}
