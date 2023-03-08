package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsDrawNineGridCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpegdi/c7fff288-63db-4521-bbe5-77e060fb0780
type TsDrawNineGridCapabilitySet struct {
	DrawNineGridSupportLevel uint32
	DrawNineGridCacheSize    uint16
	DrawNineGridCacheEntries uint16
}

func (c *TsDrawNineGridCapabilitySet) Type() uint16 {
	return CAPSTYPE_DRAWNINEGRIDCACHE
}

func (c *TsDrawNineGridCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsDrawNineGridCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
