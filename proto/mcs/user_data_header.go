package mcs

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// UserDataHeader Type
const (
	//client -> server
	CS_CORE     = 0xC001
	CS_SECURITY = 0xC002
	CS_NET      = 0xC003
	CS_CLUSTER  = 0xC004
	CS_MONITOR  = 0xC005

	//server -> client
	SC_CORE           = 0x0C01
	SC_SECURITY       = 0x0C02
	SC_NET            = 0x0C03
	SC_MCS_MSGCHANNEL = 0x0C04
	SC_MULTITRANSPORT = 0x0C08
)

// UserDataHeader
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/8a36630c-9c8e-4864-9382-2ec9d6f368ca
type UserDataHeader struct {
	Type uint16
	Len  uint16
}

func (h *UserDataHeader) Read(r io.Reader) {
	core.ReadLE(r, h)
}
