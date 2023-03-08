package mcsPdu

import (
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/x224"
	"io"
)

// ClientMcsErectDomainRequestPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/04c60697-0d9a-4afd-a0cd-2cc133151a9c
type ClientMcsErectDomainRequestPDU struct{}

func (pdu *ClientMcsErectDomainRequestPDU) Write(w io.Writer) {
	data := (&mcs.ClientErectDomain{}).Serialize()
	x224.Write(w, data)
}
