package eth

import (
	"encoding/binary"
	"errors"
	"net"
	"test/openflow/protocol/arp"
	"test/openflow/protocol/ipv4"
	"test/openflow/protocol/utils"
)

const (
	IPv4_MSG = 0x0800
	ARP_MSG  = 0x0806
	LLDP_MSG = 0x88cc
	WOL_MSG  = 0x0842
	RARP_MSG = 0x8035
	VLAN_MSG = 0x8100

	IPv6_MSG     = 0x86DD
	STP_MSG      = 0x4242
	STP_BPDU_MSG = 0xAAAA
)

type Ethernet struct {
	Delimiter uint8
	HWDst     net.HardwareAddr
	HWSrc     net.HardwareAddr
	VLANID    VLAN
	Ethertype uint16
	Data      utils.Message
}

func New() *Ethernet {
	eth := new(Ethernet)

	eth.HWDst = net.HardwareAddr(make([]byte, 6))
	eth.HWSrc = net.HardwareAddr(make([]byte, 6))
	eth.VLANID = *NewVLAN()
	eth.Ethertype = 0x800
	eth.Data = nil
	return eth
}

func (e *Ethernet) Len() (n uint16) {
	if e.VLANID.VID != 0 {
		n += 5
	}
	n += 12
	n += 2
	if e.Data != nil {
		n += e.Data.Len()
	}
	return
}

func (e *Ethernet) MarshalBinary() (data []byte, err error) {
	data = make([]byte, int(e.Len()))
	bytes := make([]byte, 0)
	n := 0
	copy(data[n:], e.HWDst)
	n += len(e.HWDst)
	copy(data[n:], e.HWSrc)
	n += len(e.HWSrc)

	if e.VLANID.VID != 0 {
		bytes, err = e.VLANID.MarshalBinary()
		if err != nil {
			return
		}
		copy(data[n:], bytes)
		n += len(bytes)
	}

	binary.BigEndian.PutUint16(data[n:n+2], e.Ethertype)
	n += 2

	if e.Data != nil {
		bytes, err = e.Data.MarshalBinary()
		if err != nil {
			return
		}
		copy(data[n:n+len(bytes)], bytes)
	}
	return
}

func (e *Ethernet) UnmarshalBinary(data []byte) error {
	if len(data) < 14 {
		return errors.New("The []byte is too short to unmarshal a full Ethernet message.")
	}

	n := 1

	e.HWDst = net.HardwareAddr(make([]byte, 6))
	e.HWSrc = net.HardwareAddr(make([]byte, 6))

	copy(e.HWDst, data[n:])
	n += len(e.HWDst)

	copy(e.HWSrc, data[n:])
	n += len(e.HWSrc)

	e.Ethertype = binary.BigEndian.Uint16(data[n:])
	if e.Ethertype == VLAN_MSG {
		e.VLANID = *new(VLAN)
		err := e.VLANID.UnmarshalBinary(data[n:])
		if err != nil {
			return err
		}
		n += int(e.VLANID.Len())

		e.Ethertype = binary.BigEndian.Uint16(data[n:])
	} else {
		e.VLANID = *new(VLAN)
		e.VLANID.VID = 0
	}
	n += 2

	switch e.Ethertype {
	case IPv4_MSG:
		e.Data = new(ipv4.IPv4)
	case ARP_MSG:
		e.Data = new(arp.ARP)
	default:
		e.Data = new(utils.Buffer)
	}
	return e.Data.UnmarshalBinary(data[n:])
}
