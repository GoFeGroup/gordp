package licPdu

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

const (
	// sent by server
	LICENSE_REQUEST    = 0x01
	PLATFORM_CHALLENGE = 0x02
	NEW_LICENSE        = 0x03
	UPGRADE_LICENSE    = 0x04

	// sent by client
	LICENSE_INFO                = 0x12
	NEW_LICENSE_REQUEST         = 0x13
	PLATFORM_CHALLENGE_RESPONSE = 0x15

	// Sent by either client or server:
	ERROR_ALERT = 0xff
)

// flags
const (
	LICENSE_PROTOCOL_VERSION_MASK = 0x0f
	EXTENDED_ERROR_MSG_SUPPORTED  = 0x80
)

// LicensingPreamble
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/73170ca2-5f82-4a2d-9d1b-b439f3d8dadc
type LicensingPreamble struct {
	BMsgType uint8
	Flags    uint8
	WMsgSize uint16 // The size in bytes of the licensing packet (including the size of the preamble).
}

func (p *LicensingPreamble) Read(r io.Reader) {
	core.ReadLE(r, p)
}

// LicenseValidClientData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/7e2fe1e0-e793-45b7-983d-af344d9ca327
type LicenseValidClientData struct {
	Preamble     LicensingPreamble
	ErrorMessage LicensingErrorMessage
}

func (d *LicenseValidClientData) Read(r io.Reader) {
	d.Preamble.Read(r)
	switch d.Preamble.BMsgType {
	case NEW_LICENSE:
		return // FIXME: OK?
	case ERROR_ALERT:
		d.ErrorMessage.Read(r)
		if d.ErrorMessage.DwErrorCode == STATUS_VALID_CLIENT &&
			d.ErrorMessage.DwStateTransaction == ST_NO_TRANSITION {
			return // FIXME: OK?
		}
		fallthrough
	case LICENSE_REQUEST:
		fallthrough
	case PLATFORM_CHALLENGE:
		fallthrough
	case UPGRADE_LICENSE:
		fallthrough
	default:
		core.Throw("not implement")
	}
}
