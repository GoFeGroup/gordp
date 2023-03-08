package licPdu

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// DwErrorCode
const (
	// Sent by client
	ERR_INVALID_SERVER_CERTIFICATE = 0x00000001
	ERR_NO_LICENSE                 = 0x00000002

	// Sent by server
	ERR_INVALID_SCOPE       = 0x00000004
	ERR_NO_LICENSE_SERVER   = 0x00000006
	STATUS_VALID_CLIENT     = 0x00000007
	ERR_INVALID_CLIENT      = 0x00000008
	ERR_INVALID_PRODUCTID   = 0x0000000B
	ERR_INVALID_MESSAGE_LEN = 0x0000000C

	// Sent by client and server
	ERR_INVALID_MAC = 0x00000003
)

// DwStateTransaction
const (
	ST_TOTAL_ABORT          = 0x00000001
	ST_NO_TRANSITION        = 0x00000002
	ST_RESET_PHASE_TO_START = 0x00000003
	ST_RESEND_LAST_MESSAGE  = 0x00000004
)

// ErrorMessage
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/f18b6c9f-f3d8-4a0e-8398-f9b153233dca
type LicensingErrorMessage struct {
	DwErrorCode        uint32
	DwStateTransaction uint32
	//BbErrorInfo        []byte
}

func (m *LicensingErrorMessage) Read(r io.Reader) {
	core.ReadLE(r, m)
}
