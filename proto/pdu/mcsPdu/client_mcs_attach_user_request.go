package mcsPdu

import (
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/x224"
	"io"
)

// ClientMcsAttachUserRequestPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/f5d6a541-9b36-4100-b78f-18710f39f247
type ClientMcsAttachUserRequestPDU struct {
}

func (pdu *ClientMcsAttachUserRequestPDU) Write(w io.Writer) {
	data := (&mcs.ClientAttachUser{}).Serialize()
	x224.Write(w, data)
}
