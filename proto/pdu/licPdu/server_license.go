package licPdu

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/sec"
	"io"
)

// ServerLicensingPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/7d941d0d-d482-41c5-b728-538faa3efb31
type ServerLicensingPDU struct {
	McsSDin                mcs.ReceiveDataResponse
	SecurityHeader         sec.TsSecurityHeader
	ValidClientLicenseData LicenseValidClientData
}

func (p *ServerLicensingPDU) Read(r io.Reader) {
	//channelId, data := p.McsSDin.Read(r)
	channelId, data := p.McsSDin.Read(r)
	core.ThrowIf(channelId != mcs.MCS_CHANNEL_GLOBAL, "invalid channel id")
	glog.Debugf("mcs read: [%v] %v - %x", channelId, len(data), data)
	r = bytes.NewReader(data)
	p.SecurityHeader.Read(r)
	core.ThrowIf(p.SecurityHeader.Flags&sec.SEC_LICENSE_PKT == 0, "invalid security header")
	p.ValidClientLicenseData.Read(r)
}
