package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	// device       string        = "vlan111"
	snapshot_len int32         = 65535
	promiscuous  bool          = false
	timeout      time.Duration = -1 * time.Second
	err          error
	handle       *pcap.Handle
)


// udp src port 53

func main() {
	device := flag.String("d", "", "device")
	bpf := flag.String("r", "", "device")
	flag.Parse()
	handle, err := pcap.OpenLive(*device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatalln(err)
	}

	defer handle.Close()

	err = handle.SetBPFFilter(*bpf)
	if err != nil {
		log.Fatalln(err)
	}

	packetSouce := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSouce.Packets() {
		fmt.Println(packet)
	}
}