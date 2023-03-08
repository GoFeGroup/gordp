package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsDrawGdiPlusCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpegdi/52635737-d144-4f47-9c88-b48ceaf3efb4
type TsDrawGdiPlusCapabilitySet struct {
	SupportLevel                uint32
	GdiPlusVersion              uint32
	CacheLevel                  uint32
	GdiPlusCacheEntries         [10]byte
	GdiPlusCacheChunkSize       [8]byte
	GdiPlusImageCacheProperties [6]byte
}

func (c *TsDrawGdiPlusCapabilitySet) Type() uint16 {
	return CAPSTYPE_DRAWGDIPLUS
}

func (c *TsDrawGdiPlusCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsDrawGdiPlusCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
