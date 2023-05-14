package tcp

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
