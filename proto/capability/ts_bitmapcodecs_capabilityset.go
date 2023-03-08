package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

type TsBitmapCodecs struct {
	BitmapCodecCount uint8
	BitmapCodecArray []byte
}

func (c *TsBitmapCodecs) Read(r io.Reader) {
	core.ReadLE(r, &c.BitmapCodecCount)
	c.BitmapCodecArray = core.ReadBytes(r, int(c.BitmapCodecCount))
}

func (c *TsBitmapCodecs) Write(w io.Writer) {
	core.WriteLE(w, &c.BitmapCodecCount)
	core.WriteFull(w, c.BitmapCodecArray)
}

// TsBitmapCodecsCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/17e80f50-d163-49de-a23b-fd6456aa472f
type TsBitmapCodecsCapabilitySet struct {
	SupportedBitmapCodecs TsBitmapCodecs // A variable-length field containing a TS_BITMAPCODECS structure (section 2.2.7.2.10.1).
}

func (c *TsBitmapCodecsCapabilitySet) Type() uint16 {
	return CAPSETTYPE_BITMAP_CODECS
}

func (c *TsBitmapCodecsCapabilitySet) Read(r io.Reader) TsCapsSet {
	c.SupportedBitmapCodecs.Read(r)
	//core.ReadLE(r,c)
	return c
}

func (c *TsBitmapCodecsCapabilitySet) Write(w io.Writer) {
	c.SupportedBitmapCodecs.Write(w)
}
