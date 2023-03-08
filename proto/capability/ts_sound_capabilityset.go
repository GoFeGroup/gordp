package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsSoundCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/fadb6a2c-18fa-4fa7-a155-e970d9b1ac59
type TsSoundCapabilitySet struct {
	Flags      uint16
	Pad2octets uint16
}

func (c *TsSoundCapabilitySet) Type() uint16 {
	return CAPSTYPE_SOUND
}

func (c *TsSoundCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsSoundCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
