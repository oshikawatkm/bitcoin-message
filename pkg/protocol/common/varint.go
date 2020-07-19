package common

import "encoding/binary"

type VarInt struct {
	Data uint64
}

func NewVarInt(u uint64) *VarInt {
	return &VarInt{
		Data: u,
	}
}

func (v *VarInt) Encode() []byte {
	if v.Data < 0xfd {
		return []byte{byte(v.Data)}
	}
	if v.Data <= 0xffff {
		b := make([]byte, 3)
		b[0] = byte(0xfd)
		binary.LittleEndian.PutUint16(b[1:], uint32(v.Data))
		return b
	}

}
