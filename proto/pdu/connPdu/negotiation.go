package connPdu

import (
	"fmt"
	"io"

	"github.com/GoFeGroup/gordp/core"
)

// Negotiation Type
const (
	TYPE_RDP_NEG_REQ     = 0x01
	TYPE_RDP_NEG_RSP     = 0x02
	TYPE_RDP_NEG_FAILURE = 0x03
)

// Negotiation Result
const (
	PROTOCOL_RDP       uint32 = 0x00000000 //Standard RDP Security https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/8e8b2cca-c1fa-456c-8ecb-a82fc60b2322
	PROTOCOL_SSL              = 0x00000001 //TLS1.0/1.1/1.2 https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/857dadbe-f01a-4047-9b63-0d5b681ad306
	PROTOCOL_HYBRID           = 0x00000002 //CredSSP https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/8e11581d-094f-461a-9fde-ba51af90cf8b
	PROTOCOL_RDSTLS           = 0x00000004 //https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/83d1186d-cab6-4ad8-8c5f-203f95e192aa
	PROTOCOL_HYBRID_EX        = 0x00000008 //https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/d0e560a3-25cb-4563-8bdc-6c4cc625bbfc
	PROTOCOL_RDSAAD           = 0x00000010 //https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/dc43f040-d75d-49a9-90c6-0c9999281136
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
	core.ThrowIf(nego.Type != TYPE_RDP_NEG_RSP, fmt.Errorf("invalid nego type: %v", nego.Type))
	core.ThrowIf(nego.Length != 8, fmt.Errorf("invalid nego.length: %v", nego.Length))
}
