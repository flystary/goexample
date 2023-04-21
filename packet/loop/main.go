package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device       string        = "enp0s8"
	snapshot_len int32         = 65535
	promiscuous  bool          = false
	timeout      time.Duration = -1 * time.Second
	err          error
	handle       *pcap.Handle
)


func main() {
	handle, err =pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatalln(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println("----packet-Layers-----")
		for _, layer := range packet.Layers() {
			fmt.Println(layer.LayerType())
		}
		fmt.Println("------------------")

		// IPv4
		ip4Layer := packet.Layer(layers.LayerTypeIPv4)
		if ip4Layer != nil {
			fmt.Println("IPv4 layer detected.")
			ip, _ := ip4Layer.(*layers.IPv4)
			fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
			fmt.Println("Protocol: ", ip.Protocol)
			fmt.Println()
		} else {
			fmt.Println("No IPv4 layer detected.")
		}

		// TCP
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			fmt.Println("TCP layer detected.")
			tcp, _ := tcpLayer.(*layers.TCP)
			fmt.Println("ACK: ", tcp.ACK)
			fmt.Println("SYN: ", tcp.SYN)
			fmt.Println("Seq: ", tcp.Seq)
			fmt.Println("DstPort: ", tcp.DstPort)
			fmt.Println("SrcPort: ", tcp.SrcPort)
			fmt.Println()
		} else {
			fmt.Println("No TCP layer detected.")
		}
	}
}