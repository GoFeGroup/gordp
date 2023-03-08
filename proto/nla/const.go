package nla

import "bytes"

const (
	WINDOWS_MINOR_VERSION_0 = 0x00
	WINDOWS_MINOR_VERSION_1 = 0x01
	WINDOWS_MINOR_VERSION_2 = 0x02
	WINDOWS_MINOR_VERSION_3 = 0x03

	WINDOWS_MAJOR_VERSION_5 = 0x05
	WINDOWS_MAJOR_VERSION_6 = 0x06
	NTLMSSP_REVISION_W2K3   = 0x0F
)

type Field struct {
	Len    uint16
	MaxLen uint16
	Offset uint32
}

func (field *Field) Set(length uint16, offset uint32) {
	field.Len = length
	field.MaxLen = length
	field.Offset = offset
}

type msgBase struct{}

func (m *msgBase) GetField(data []byte, offset uint32, field *Field) []byte {
	if field.Len == 0 {
		return make([]byte, 0)
	}
	start := field.Offset - offset
	return data[start : start+uint32(field.Len)]
}

func concat(bs ...[]byte) []byte {
	return bytes.Join(bs, nil)
}
