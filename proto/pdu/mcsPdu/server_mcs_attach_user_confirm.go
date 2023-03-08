package mcsPdu

import (
	"bytes"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/x224"
	"io"
)

// ServerMcsAttachUserConfirmPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/3b3d850b-99b1-4a9a-852b-1eb2da5024e5
type ServerMcsAttachUserConfirmPDU struct {
	McsAUcf mcs.ServerAttachUserConfirm
}

func (pdu *ServerMcsAttachUserConfirmPDU) Read(r io.Reader) {
	data := x224.Read(r)
	glog.Debugf("receive attach user confirm: %v - %x", len(data), data)
	pdu.McsAUcf.Read(bytes.NewReader(data))
}
