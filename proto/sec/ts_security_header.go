package sec

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

/* Security Header Flags */
const (
	SEC_EXCHANGE_PKT       = 0x0001
	SEC_TRANSPORT_REQ      = 0x0002
	SEC_TRANSPORT_RSP      = 0x0004
	SEC_ENCRYPT            = 0x0008
	SEC_RESET_SEQNO        = 0x0010
	SEC_IGNORE_SEQNO       = 0x0020
	SEC_INFO_PKT           = 0x0040
	SEC_LICENSE_PKT        = 0x0080
	SEC_LICENSE_ENCRYPT_CS = 0x0200
	SEC_LICENSE_ENCRYPT_SC = 0x0200
	SEC_REDIRECTION_PKT    = 0x0400
	SEC_SECURE_CHECKSUM    = 0x0800
	SEC_AUTODETECT_REQ     = 0x1000
	SEC_AUTODETECT_RSP     = 0x2000
	SEC_HEARTBEAT          = 0x4000
	SEC_FLAGSHI_VALID      = 0x8000
)

// TsSecurityHeader
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/e13405c5-668b-4716-94b2-1c2654ca1ad4
type TsSecurityHeader struct {
	Flags   uint16
	FlagsHi uint16 // unused
}

func (h *TsSecurityHeader) Write(w io.Writer) {
	core.WriteLE(w, h)
}

func (h *TsSecurityHeader) Read(r io.Reader) {
	core.ReadLE(r, h)
	glog.Debugf("security header: %+v", h)
}

func NewTsSecurityHeader(flags uint16) *TsSecurityHeader {
	return &TsSecurityHeader{Flags: flags}
}
