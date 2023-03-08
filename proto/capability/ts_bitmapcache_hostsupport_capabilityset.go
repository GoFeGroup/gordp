package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsBitmapCacheHostSupportCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/fc05c385-46c3-42cb-9ed2-c475a3990e0b
type TsBitmapCacheHostSupportCapabilitySet struct {
	CacheVersion uint8
	Pad1         uint8
	Pad2         uint16
}

func (c *TsBitmapCacheHostSupportCapabilitySet) Type() uint16 {
	return CAPSTYPE_BITMAPCACHE_HOSTSUPPORT
}

func (c *TsBitmapCacheHostSupportCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsBitmapCacheHostSupportCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
