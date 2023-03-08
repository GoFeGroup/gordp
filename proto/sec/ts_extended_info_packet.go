package sec

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

const (
	ADDRESS_FAMILY_INET  uint16 = 0x0002
	ADDRESS_FAMILY_INET6        = 0x0017
)

// TsExtendedInfoPacket
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/05ada9e4-a468-494b-8694-eb806a0ecc89
type TsExtendedInfoPacket struct {
	ClientAddressFamily uint16
	CbClientAddress     uint16
	ClientAddress       []byte
	CbClientDir         uint16
	ClientDir           []byte
	ClientTimeZone      []byte
	ClientSessionId     uint32
	PerformanceFlags    uint32
	AutoReconnect       *TsAutoReconnectInfo
}

func (p *TsExtendedInfoPacket) Write(w io.Writer) {
	core.WriteLE(w, p.ClientAddressFamily)
	core.WriteLE(w, p.CbClientAddress)
	core.WriteFull(w, p.ClientAddress)
	core.WriteLE(w, p.CbClientDir)
	core.WriteFull(w, p.ClientDir)
	core.WriteFull(w, p.ClientTimeZone)
	core.WriteLE(w, p.ClientSessionId)
	core.WriteLE(w, p.PerformanceFlags)
	if p.AutoReconnect != nil {
		p.AutoReconnect.Write(w)
	}
}

func NewExtendedInfoPacket() *TsExtendedInfoPacket {
	return &TsExtendedInfoPacket{
		ClientAddressFamily: ADDRESS_FAMILY_INET,
		CbClientAddress:     2,
		ClientAddress:       []byte{0, 0},
		CbClientDir:         2,
		ClientDir:           []byte{0, 0},
		ClientTimeZone:      make([]byte, 172),
		ClientSessionId:     0,
		PerformanceFlags:    0,
		AutoReconnect:       nil,
	}
}
