package connPdu

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// Negotiation Type
const (
	TYPE_RDP_NEG_REQ     = 0x01
	TYPE_RDP_NEG_RSP     = 0x02
	TYPE_RDP_NEG_FAILURE = 0x03
)

// Negotiation Result
const (
	PROTOCOL_RDP       uint32 = 0x00000000
	PROTOCOL_SSL              = 0x00000001
	PROTOCOL_HYBRID           = 0x00000002
	PROTOCOL_HYBRID_EX        = 0x00000008
)

// Negotiation
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/b2975bdc-6d56-49ee-9c57-f2ff3a0b6817
type Negotiation struct {
	Type   uint8
	Flag   uint8
	Length uint16
	Result uint32
}

func (nego *Negotiation) Read(r io.Reader) {
	core.ReadLE(r, nego)
}
