package ipv4

import (
	"net"
	"test/openflow/protocol/utils"
)

const (
	Type_ICMP     = 0x01
	Type_TCP      = 0x06
	Type_UDP      = 0x11
	Type_IPv6     = 0x29
	Type_IPv6ICMP = 0x3a
)

type IPv4 struct {
	Version        uint8 //4-bits
	IHL            uint8 //4-bits
	DSCP           uint8 //6-bits
	ECN            uint8 //2-bits
	Length         uint16
	Id             uint16
	Flags          uint16 //3-bits
	FragmentOffset uint16 //13-bits
	TTL            uint8
	Protocol       uint8
	Checksum       uint16
	NWSrc          net.IP
	NWDst          net.IP
	Options        utils.Buffer
	Data           utils.Message
}

func New() *IPv4 {
	ip := new(IPv4)
	ip.NWSrc = make([]byte, 4)
	ip.NWDst = make([]byte, 4)
	ip.Options = *new(utils.Buffer)

	return ip
}
