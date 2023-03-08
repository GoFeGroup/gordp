package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"io"
)

// TsBitmapCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/76670547-e35c-4b95-a242-5729a21b83f6
type TsBitmapCapabilitySet struct {
	PreferredBitsPerPixel    uint16
	Receive1BitPerPixel      uint16 // This field is ignored and SHOULD be set to TRUE (0x0001)
	Receive4BitsPerPixel     uint16 // This field is ignored and SHOULD be set to TRUE (0x0001)
	Receive8BitsPerPixel     uint16 // This field is ignored and SHOULD be set to TRUE (0x0001)
	DesktopWidth             uint16
	DesktopHeight            uint16
	Pad2octets               uint16
	DesktopResizeFlag        uint16
	BitmapCompressionFlag    uint16 // This field MUST be set to TRUE (0x0001) because support for compressed bitmaps is required for a connection to proceed.
	HighColorFlags           uint8  // This field is ignored and SHOULD be set to zero.
	DrawingFlags             uint8
	MultipleRectangleSupport uint16
	Pad2octetsB              uint16
}

func (c *TsBitmapCapabilitySet) Type() uint16 {
	return CAPSTYPE_BITMAP
}

func (c *TsBitmapCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsBitmapCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}

func NewTsBitmapCapabilitySet() *TsBitmapCapabilitySet {
	return &TsBitmapCapabilitySet{
		Receive1BitPerPixel:      0x0001,
		Receive4BitsPerPixel:     0x0001,
		Receive8BitsPerPixel:     0x0001,
		BitmapCompressionFlag:    0x0001,
		MultipleRectangleSupport: 0x0001,
		DesktopWidth:             1280,
		DesktopHeight:            800,
		PreferredBitsPerPixel:    mcs.HIGH_COLOR_24BPP,
		DesktopResizeFlag:        0x0001, // support Resize
	}
}
