package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"github.com/robfig/cron"
	"net"
	"os"
	"time"
)

type Packets struct {
	Time     string        `json:"time"`
	Addr     string        `json:"addr"`
	Ip       *net.IPAddr   `json:"ip""`
	Sent     int           `json:"sent"`
	Received int           `json:"received"`
	Loss     float64       `json:"loss"`
	Avg      time.Duration `json:"avg""`
}

func NetPing(ip string) Packets {
	var ipInfo Packets
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(true)
	pinger.Interval = time.Duration(500) * time.Millisecond
	pinger.Count = 10
	//pinger.Size = 10
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	ipInfo = Packets{
		Time:     time.Now().Format("2006/01/02 15:04:05"),
		Addr:     pinger.Addr(),
		Ip:       pinger.IPAddr(),
		Sent:     pinger.PacketsSent,
		Received: pinger.PacketsRecv,
		Loss:     stats.PacketLoss,
		Avg:      stats.AvgRtt,
	}

	fmt.Printf("-------------------链路测试[%s]--------------------\n", ipInfo.Ip)
	//fmt.Println(ret.Ip)
	fmt.Printf(" 时间:%s", ipInfo.Time)
	fmt.Printf(" 已接收:%d", ipInfo.Received)
	fmt.Printf(" 时延:%v", ipInfo.Avg)
	fmt.Printf(" 丢包率:%v\n", ipInfo.Loss)
	fmt.Println("----------------------------><-------------------------------")

	return ipInfo
}

func touchTxt(ipInfo Packets) {
	ret := ipInfo
	filePtr, err := os.OpenFile("./netping.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("创建文件失败，err=", err)
		return
	}
	//defer filePtr.Close()
	_, err = fmt.Fprintln(filePtr, ret)
	if err != nil {
		panic(err)
		return
	}
	return
}

func CreateCron(ip string) {
	c := cron.New()
	spec := "1 * * * * ?"

	if err := c.AddFunc(spec, func() {
		//pingIP
		ret := NetPing(ip)
		//创建文件
		touchTxt(ret)
	}); err != nil {
		panic(err)
	}
	//运行定时任务
	c.Start()
}

func main() {
	var ip string

	fmt.Print("请输入测试IP或域名: ")
	_, err := fmt.Scanln(&ip)
	if err != nil {
		panic(err)
	}

	//1分钟执行一次
	go CreateCron(ip)
	for {
	}
}
