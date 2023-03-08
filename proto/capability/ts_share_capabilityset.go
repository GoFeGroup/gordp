package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsShareCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/75caa232-1929-41bb-9d59-6f8aad59ecf5
type TsShareCapabilitySet struct {
	NodeId     uint16
	Pad2octets uint16
}

func (c *TsShareCapabilitySet) Type() uint16 {
	return CAPSTYPE_SHARE
}

func (c *TsShareCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsShareCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
