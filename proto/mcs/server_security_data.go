package mcs

import (
	"bytes"
	"io"

	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
)

// encryptionMethod
const (
	ENCRYPTION_METHOD_NONE   = 0x00000000
	ENCRYPTION_METHOD_40BIT  = 0x00000001
	ENCRYPTION_METHOD_128BIT = 0x00000002
	ENCRYPTION_METHOD_56BIT  = 0x00000008
	ENCRYPTION_METHOD_FIPS   = 0x00000010
)

// encryptionLevel
const (
	ENCRYPTION_LEVEL_NONE              = 0x00000000
	ENCRYPTION_LEVEL_LOW               = 0x00000001
	ENCRYPTION_LEVEL_CLIENT_COMPATIBLE = 0x00000002
	ENCRYPTION_LEVEL_HIGH              = 0x00000003
	ENCRYPTION_LEVEL_FIPS              = 0x00000004
)

// ServerSecurityData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/3e86b68d-3e2e-4433-b486-878875778f4b
type ServerSecurityData struct {
	EncryptionMethod  uint32
	EncryptionLevel   uint32
	ServerRandomLen   uint32 //0x00000020
	ServerCertLen     uint32
	ServerRandom      []byte
	ServerCertificate ServerCertificate
}

func (d *ServerSecurityData) Read(r io.Reader) {
	defer func() { glog.Debugf("server security data: %+v", d) }()

	core.ReadLE(r, &d.EncryptionMethod)
	core.ReadLE(r, &d.EncryptionLevel)

	// Win10: 0-0
	// WinXP: 2-2
	glog.Debugf("%v-%v", d.EncryptionMethod, d.EncryptionLevel)

	if d.EncryptionMethod == 0 && d.EncryptionLevel == 0 {
		return
	}

	core.ReadLE(r, &d.ServerRandomLen)
	core.ReadLE(r, &d.ServerCertLen)
	d.ServerRandom = core.ReadBytes(r, int(d.ServerRandomLen))
	serverCertData := core.ReadBytes(r, int(d.ServerCertLen))
	d.ServerCertificate.Read(bytes.NewReader(serverCertData))
}
