package main

import (
	"fmt"
	"time"
)

func main() {
	//timer1()
	//timer2()
	//timer3()
	//timer4()
	timer5()

}
func timer1() {
	// 1.timer基本使用
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2.Format("2006-01-02 15:04:05"))
}

func timer2() {
	//2. 验证只能响应一次
	timer2 := time.NewTimer(time.Second * 2)
	for {
		<-timer2.C
		fmt.Println("时间到")
	}
}

// 3.timer实现延时的功能
func timer3() {
	time.Sleep(time.Second)
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	<-time.After(2 * time.Second)
	fmt.Println("2秒到")
}

// 4.停止定时器
func timer4() {
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行了")
	}()

	b := timer4.Stop()
	if b {
		fmt.Println("timer4已关闭")
	}
}

// 5.重置定时器
func timer5() {
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

	for {
	}
}
