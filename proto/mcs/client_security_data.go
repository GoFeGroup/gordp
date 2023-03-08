package mcs

import (
	"github.com/GoFeGroup/gordp/core"
)

// ClientSecurityData EncryptionMethods
const (
	ENCRYPTION_FLAG_40BIT  uint32 = 0x00000001
	ENCRYPTION_FLAG_128BIT        = 0x00000002
	ENCRYPTION_FLAG_56BIT         = 0x00000008
	FIPS_ENCRYPTION_FLAG          = 0x00000010
)

// ClientSecurityData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/6b58e11e-a32b-4903-b736-339f3cfe46ec
type ClientSecurityData struct {
	Header               UserDataHeader // CS_SECURITY
	EncryptionMethods    uint32
	ExtEncryptionMethods uint32 // only for French locale client, otherwise must be set to zero
}

func NewClientSecurityData() *ClientSecurityData {
	return &ClientSecurityData{
		Header:               UserDataHeader{Type: CS_SECURITY, Len: 0x0C},
		EncryptionMethods:    ENCRYPTION_FLAG_40BIT | ENCRYPTION_FLAG_56BIT | ENCRYPTION_FLAG_128BIT,
		ExtEncryptionMethods: 00,
	}
}

func (securityData *ClientSecurityData) Serialize() []byte {
	return core.ToLE(securityData)
}
