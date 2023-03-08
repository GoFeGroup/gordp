package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsColorTableCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpegdi/2b7c6946-3612-4291-95a8-03b7b1387eaf
type TsColorTableCapabilitySet struct {
	CacheSize  uint16
	Pad2octets uint16
}

func (c *TsColorTableCapabilitySet) Type() uint16 {
	return CAPSTYPE_COLORCACHE
}

func (c *TsColorTableCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsColorTableCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
