package mcs

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
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
	defer func() { glog.Debugf("server security data: %0#x", d) }()

	core.ReadLE(r, &d.EncryptionMethod)
	core.ReadLE(r, &d.EncryptionLevel)
	if d.EncryptionMethod == 0 && d.EncryptionLevel == 0 {
		return
	}

	core.ReadLE(r, &d.ServerRandomLen)
	core.ReadLE(r, &d.ServerCertLen)
	d.ServerRandom = make([]byte, d.ServerRandomLen)
	_, err := io.ReadFull(r, d.ServerRandom)
	core.ThrowError(err)

	// read certdata
	data := make([]byte, d.ServerCertLen)
	_, err = io.ReadFull(r, data)
	core.ThrowError(err)

	d.ServerCertificate.Read(bytes.NewReader(data))
}
