package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsLargePointerCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/41323437-c753-460e-8108-495a6fdd68a8
type TsLargePointerCapabilitySet struct {
	SupportFlags uint16
}

func (c *TsLargePointerCapabilitySet) Type() uint16 {
	return CAPSETTYPE_LARGE_POINTER
}

func (c *TsLargePointerCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsLargePointerCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
