package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-ping/ping"
	"os"
	"time"
)

func NetPing(ip string) map[string]interface{} {
	var ipInfo = make(map[string]interface{})
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(true)
	pinger.Count = 10
	//pinger.Size = 10
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()

	ipInfo["Addr"] = pinger.Addr()
	ipInfo["Ip"] = pinger.IPAddr()
	ipInfo["Sent"] = pinger.PacketsSent
	ipInfo["Received"] = pinger.PacketsRecv
	ipInfo["Loss"] = stats.PacketLoss
	ipInfo["Avg"] = stats.AvgRtt
	ipInfo["Time"] = time.Now().Format("2006/01/02 15:04:05")

	fmt.Printf("-------------------链路测试[%s]--------------------\n", ipInfo["Ip"])
	//fmt.Println(ret.Ip)
	fmt.Printf(" 时间:%s", ipInfo["Time"])
	fmt.Printf(" 已接收:%d", ipInfo["Received"])
	fmt.Printf(" 时延:%v", ipInfo["Avg"])
	fmt.Printf(" 丢包率:%v\n", ipInfo["Loss"])
	fmt.Println("----------------------------><-------------------------------")

	return ipInfo
}

func main() {
	var ip string
	fmt.Print("请输入测试的链路IP: ")
	_, err := fmt.Scanln(&ip)
	if err != nil {
		return
	}
	for {
		ret := NetPing(ip)
		filePtr, err := os.OpenFile("./netping.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("创建文件失败，err=", err)
			return
		}
		//defer filePtr.Close()
		encoder := json.NewEncoder(filePtr)

		//将实例编码到文件中
		err = encoder.Encode(ret)

		if err != nil {
			fmt.Println("编码失败，err=", err)
		}
	}

}
