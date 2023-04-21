package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//if len(os.Args) > 0 {
	//	for index, arg := range os.Args {
	//		fmt.Printf("args[%d]=%v\n", index, arg)
	//	}
	//}
	var name string
	var retime string

	flag.StringVar(&name, "name", "许捧", "姓名")
	flag.StringVar(&retime, "rt", "2022/01/01", "离职时间")

	flag.Parse()
	updatetime := lizhiTime(retime)
	fmt.Printf("亲爱的的 %v 同学\n", name)
	fmt.Printf("距离您离开达科还有:%v天\n", updatetime) //int(updatetime.Hours())/24

}

func lizhiTime(retime string) int {
	now := time.Now()

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed, err:%v\n", err)
	}
	timeObj, err := time.ParseInLocation("2006/01/02", retime, loc)
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
	}
	td := int(timeObj.Sub(now).Hours() / 24)

	return td
}
