package main

import (
	"fmt"
	"unsafe"
)

type TCP struct {
	PortSrc uint16 // 2
	PortDst uint16 // 2
	SeqNum  uint32 // 4
	AckNum  uint32 // 4

	WinSize  uint16 // 2
	Checksum uint16 // 2
	UrgFlag  uint16 // 2
	Data     []byte // 24
}

func main() {
	tcp := TCP{}
	fmt.Printf("%v\n", unsafe.Sizeof(TCP{}))
	fmt.Println()
	fmt.Printf("%v\n", unsafe.Sizeof(tcp.PortSrc))
	fmt.Printf("%v\n", unsafe.Sizeof(tcp.PortDst))
	fmt.Printf("%v\n", unsafe.Sizeof(tcp.SeqNum))
	fmt.Printf("%v\n", unsafe.Sizeof(tcp.AckNum))
	fmt.Printf("%v\n", unsafe.Sizeof(tcp.WinSize))
	fmt.Printf("%v\n", unsafe.Sizeof(tcp.Checksum))
	fmt.Printf("%v\n", unsafe.Sizeof(tcp.UrgFlag))
	fmt.Printf("%v\n", unsafe.Sizeof(tcp.Data))
	fmt.Println()
	fmt.Printf("%v\n", unsafe.Sizeof([]byte{}))
}
