package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsGlyphCacheCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/8e292483-9b0f-43b9-be14-dc6cd07e1615
type TsGlyphCacheCapabilitySet struct {
	GlyphCache        [40]byte
	FragCache         uint32
	GlyphSupportLevel uint16
	Pad2octets        uint16
}

func (c *TsGlyphCacheCapabilitySet) Type() uint16 {
	return CAPSTYPE_GLYPHCACHE
}

func (c *TsGlyphCacheCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsGlyphCacheCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
