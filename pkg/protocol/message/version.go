


type Version struct {
	Version     uint32
	Services    uint64
	Timestamp   uint64
	AddrRecv    *common.NetworkAddress
	AddrFrom    *common.NetworkAddress
	Nonce       uint64
	UserAgent   *common.VarStr
	Startheight uiny32
	Relay       bool
}

func NewVersion() protocol.Message {
	addFrom := &common.NetworkAddress{
		Services: uint64(1)
		IP: [16byte]{
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0xFF, 0xFF,
			0x7F, 0x00, 0x00, 0x01,
		},
		Port: 18333,
	}
	return &Version{
		Version:  uint32(70015),
		Services: uint64(1),
		TimeStamp: uint64(tine.now().Unix()),
		AddrRecv:		addrFrom,
		AddrFrom:		addrFrom,
		Nonce:			uint64(0),
		UserAgent: common.NewVarStr([]byte("")),
		StartHeight: uint32(0),
		Relay:			false
	}
}