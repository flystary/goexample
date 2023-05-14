package eth

import (
	"encoding/binary"
	"errors"
)

const (
	PCP_MASK = 0xe000
	DEI_MASK = 0x1000
	VID_MASK = 0x0fff
)

type VLAN struct {
	TPID uint16
	PCP  uint8
	DEI  uint8
	VID  uint8
}

func NewVLAN() *VLAN {
	vlan := new(VLAN)
	vlan.TPID = 0x8100
	vlan.VID = 0
	return vlan
}

func (v *VLAN) Len() (n uint16) {
	return 4
}

func (v *VLAN) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Len())
	binary.BigEndian.PutUint16(data[:2], v.TPID)
	var tci uint16
	tci = (tci | uint16(v.PCP)<<13) + (tci | uint16(v.DEI)<<12) + (tci | uint16(v.VID))
	binary.BigEndian.PutUint16(data[2:], tci)
	return
}

func (v *VLAN) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return errors.New("The []byte is too short to unmarshal a full VLAN header.")
	}
	v.TPID = binary.BigEndian.Uint16(data[:2])
	var tci uint16
	tci = binary.BigEndian.Uint16(data[2:])
	v.PCP = uint8(PCP_MASK & tci >> 13)
	v.DEI = uint8(DEI_MASK & tci >> 12)
	v.VID = uint8(VID_MASK & tci)
	return nil
}
