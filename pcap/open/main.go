package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)
var (
	handle *pcap.Handle
	err    error
)

func main() {
	handle, err = pcap.OpenOffline("./test.pcap")
	if err != nil {
		log.Fatalln(err)
	}
	defer handle.Close()

	// Process packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}