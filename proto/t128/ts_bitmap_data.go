package t128

import (
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

const (
	BITMAP_COMPRESSION        = 0x0001
	NO_BITMAP_COMPRESSION_HDR = 0x0400
)

// TsBitmapData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/84a3d4d2-5523-4e49-9a48-33952c559485
type TsBitmapData struct {
	DestLeft         uint16
	DestTop          uint16
	DestRight        uint16
	DestBottom       uint16
	Width            uint16
	Height           uint16
	BitsPerPixel     uint16
	Flags            uint16
	BitmapLength     uint16
	BitmapComprHdr   *TsCdHeader
	BitmapDataStream []byte
}

func (d *TsBitmapData) Read(r io.Reader) {
	core.ReadLE(r, &d.DestLeft)
	core.ReadLE(r, &d.DestTop)
	core.ReadLE(r, &d.DestRight)
	core.ReadLE(r, &d.DestBottom)
	core.ReadLE(r, &d.Width)
	core.ReadLE(r, &d.Height)
	core.ReadLE(r, &d.BitsPerPixel)
	core.ReadLE(r, &d.Flags)
	core.ReadLE(r, &d.BitmapLength)

	core.ThrowIf(d.Width == 0 || d.Height == 0,
		fmt.Errorf("invalid BITMAP_DATA: width=%v, height=%v", d.Width, d.Height))
	if d.Flags&BITMAP_COMPRESSION != 0 {
		if d.Flags&NO_BITMAP_COMPRESSION_HDR == 0 {
			d.BitmapComprHdr = (&TsCdHeader{}).Read(r)
			glog.Debugf("compressionHeader: %+v", d.BitmapComprHdr)
		}
		//glog.Debugf("[!] compression = true")
	}
	if d.BitmapLength > 0 {
		d.BitmapDataStream = core.ReadBytes(r, int(d.BitmapLength))
	}
	//glog.Debugf("bitmap: %x", d.BitmapDataStream)
	glog.Debugf("[%v,%v,%v,%v],bpp: %v, len: %v", d.DestLeft, d.DestTop, d.Width, d.Height, d.BitsPerPixel, d.BitmapLength)
	//core.ShowHex(d.BitmapDataStream)
	//if len(d.BitmapDataStream) > 64 {
	//	core.ShowHex(d.BitmapDataStream[:64])
	//} else {
	//	core.ShowHex(d.BitmapDataStream)
	//}

	// Just For Test
	//if d.Flags&BITMAP_COMPRESSION != 0 && d.BitsPerPixel == 32 { // RDP6
	//	bmp := bitmap.NewBitMapFromRDP6(d.BitmapDataStream)
	//	glog.Debugf("png: %x", bmp.ToPng())
	//} else if d.Flags&BITMAP_COMPRESSION != 0 { // RLE
	//	bmp := bitmap.NewBitmapFromRLE(d.Width, d.Height, d.BitsPerPixel, d.BitmapDataStream)
	//	glog.Debugf("png: %x", bmp.ToPng())
	//}
}
