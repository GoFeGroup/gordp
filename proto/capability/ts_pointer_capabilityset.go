package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsPointerCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/925e2c05-c13f-44b1-aa20-23082051fef9
type TsPointerCapabilitySet struct {
	ColorPointerFlag      uint16
	ColorPointerCacheSize uint16
	//PointerCacheSize      uint16 // only server need
}

func (c *TsPointerCapabilitySet) Type() uint16 {
	return CAPSTYPE_POINTER
}

func (c *TsPointerCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsPointerCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
