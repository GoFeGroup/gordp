package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsOffscreenCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/412fa921-2faa-4f1b-ab5f-242cdabc04f9
type TsOffscreenCapabilitySet struct {
	SupportLevel uint32
	CacheSize    uint16
	CacheEntries uint16
}

func (c *TsOffscreenCapabilitySet) Type() uint16 {
	return CAPSTYPE_OFFSCREENCACHE
}

func (c *TsOffscreenCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsOffscreenCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
