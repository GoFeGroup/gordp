package t128

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

const (
	PDUTYPE_DEMANDACTIVEPDU  = 0x11 //Demand Active PDU (section 2.2.1.13.1).
	PDUTYPE_CONFIRMACTIVEPDU = 0x13 //Confirm Active PDU (section 2.2.1.13.2).
	PDUTYPE_DEACTIVATEALLPDU = 0x16 //Deactivate All PDU (section 2.2.3.1).
	PDUTYPE_DATAPDU          = 0x17 //Data PDU (actual type is revealed by the pduType2 field in the Share Data Header (section 2.2.8.1.1.1.2) structure).
	PDUTYPE_SERVER_REDIR_PKT = 0x1A //Enhanced Security Server Redirection PDU (section 2.2.13.3.1).
)

// TsShareControlHeader
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/73d01865-2eae-407f-9b2c-87e31daac471
type TsShareControlHeader struct {
	TotalLength uint16
	PDUType     uint16
	PDUSource   uint16
}

func (h *TsShareControlHeader) Read(r io.Reader) {
	core.ReadLE(r, h)
	glog.Debugf("share control header: %+v", h)
	//h.PDUType &= 0xF // Copy from FreeRDP C.rdp_read_share_control_header
}

func (h *TsShareControlHeader) Write(w io.Writer) {
	core.WriteLE(w, h)
}

func (h *TsShareControlHeader) Serialize() []byte {
	buff := new(bytes.Buffer)
	h.Write(buff)
	glog.Debugf("share control header: %+v", h)
	return buff.Bytes()
}
