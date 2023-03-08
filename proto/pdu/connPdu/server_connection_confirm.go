package connPdu

import (
	"bytes"
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/x224"
	"io"
)

// ServerConnectionConfirmPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/13757f8f-66db-4273-9d2c-385c33b1e483
type ServerConnectionConfirmPDU struct {
	ProtocolNeg Negotiation
}

func (pdu *ServerConnectionConfirmPDU) Read(r io.Reader) {
	typ, data := x224.ReadConfirm(r)
	if typ != x224.TPDU_CONNECTION_CONFIRM {
		core.ThrowError(fmt.Errorf("invalid response flag: %x, should be TYPE_RDP_NEG_RSP", typ))
	}
	pdu.ProtocolNeg.Read(bytes.NewReader(data))
}
