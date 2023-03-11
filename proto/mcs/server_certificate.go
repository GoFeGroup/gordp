package mcs

import (
	"io"

	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
)

const (
	CERT_CHAIN_VERSION_1 = 0x00000001
	CERT_CHAIN_VERSION_2 = 0x00000002
)

type CertData interface {
	GetPublicKey() (uint32, []byte)
	Verify() bool
	Read(io.Reader)
}

// ServerCertificate
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/54e72cc6-3422-404c-a6b4-2486db125342
type ServerCertificate struct {
	DwVersion uint32
	CertData  CertData
}

func (c *ServerCertificate) Read(r io.Reader) {
	core.ReadLE(r, &c.DwVersion)

	glog.Debugf("dwVersion: %v", c.DwVersion&0x7fffffff)
	switch c.DwVersion & 0x7fffffff {
	case CERT_CHAIN_VERSION_1:
		c.CertData = &ProprietaryServerCertificate{}
	case CERT_CHAIN_VERSION_2:
		c.CertData = &X509CertificateChain{}
	default:
		core.Throw("invalid certificate version")
	}
	c.CertData.Read(r)
}
