package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsBrushCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/8b6a830f-3dde-4a84-9250-21ffa7d2e342
type TsBrushCapabilitySet struct {
	SupportLevel uint32
}

func (c *TsBrushCapabilitySet) Type() uint16 {
	return CAPSTYPE_BRUSH
}

func (c *TsBrushCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsBrushCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
