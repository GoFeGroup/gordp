package mcs

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// ServerCoreData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/379a020e-9925-4b4f-98f3-7d634e10b411
type ServerCoreData struct {
	Version                  uint32
	ClientRequestedProtocols uint32
	EarlyCapabilityFlags     uint32
}

func (d *ServerCoreData) Read(r io.Reader) {
	core.ReadLE(r, d)
}
