package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket/pcap"
)


func main() {
	var devices []pcap.Interface
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}

	for _, device := range devices {
		name := device.Name
		if !strings.Contains(name, "enp") && !strings.Contains(name, "eth") {
			continue
		}
		fmt.Println("Name: ", device.Name)
		// fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ")
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}