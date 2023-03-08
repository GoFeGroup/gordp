package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsFontCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/18b4ccdc-e5b0-43c4-a453-cfa8c9feb2a4
type TsFontCapabilitySet struct {
	SupportFlags uint16
	Pad2octets   uint16
}

func (c *TsFontCapabilitySet) Type() uint16 {
	return CAPSTYPE_FONT
}

func (c *TsFontCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsFontCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
