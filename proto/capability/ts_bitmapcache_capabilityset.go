package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsBitmapCacheCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/101d40a7-56c0-40e1-bcb9-1475ff63cb9d
type TsBitmapCacheCapabilitySet struct {
	Pad1                  uint32
	Pad2                  uint32
	Pad3                  uint32
	Pad4                  uint32
	Pad5                  uint32
	Pad6                  uint32
	Cache0Entries         uint16
	Cache0MaximumCellSize uint16
	Cache1Entries         uint16
	Cache1MaximumCellSize uint16
	Cache2Entries         uint16
	Cache2MaximumCellSize uint16
}

func (c *TsBitmapCacheCapabilitySet) Type() uint16 {
	return CAPSTYPE_BITMAPCACHE
}

func (c *TsBitmapCacheCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsBitmapCacheCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
