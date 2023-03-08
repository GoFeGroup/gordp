package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// WindowListCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdperp/82ec7a69-f7e3-4294-830d-666178b35d15
type WindowListCapabilitySet struct {
	WndSupportLevel     uint32
	NumIconCaches       uint8
	NumIconCacheEntries uint16
}

func (c *WindowListCapabilitySet) Type() uint16 {
	return CAPSTYPE_WINDOW
}

func (c *WindowListCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *WindowListCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
