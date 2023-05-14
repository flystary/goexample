package eth

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestVlan(t *testing.T) {
	vlan := VLAN{}
	fmt.Printf("%d\n", unsafe.Sizeof(vlan))
}
