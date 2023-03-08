package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsBitmapCacheCapabilitySetRev2
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/a5b9b9a6-5f67-4089-a95d-009bc8e25bfc
type TsBitmapCacheCapabilitySetRev2 struct {
	CacheFlags           uint16
	Pad2                 uint8
	NumCellCaches        uint8
	BitmapCache0CellInfo uint32
	BitmapCache1CellInfo uint32
	BitmapCache2CellInfo uint32
	BitmapCache3CellInfo uint32
	BitmapCache4CellInfo uint32
	Pad3                 [12]byte
}

func (c *TsBitmapCacheCapabilitySetRev2) Type() uint16 {
	return CAPSTYPE_BITMAPCACHE_REV2
}

func (c *TsBitmapCacheCapabilitySetRev2) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsBitmapCacheCapabilitySetRev2) Write(w io.Writer) {
	core.WriteLE(w, c)
}
