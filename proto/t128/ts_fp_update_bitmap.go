package t128

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

// TsFpUpdateBitmap
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/d681bb11-f3b5-4add-b092-19fe7075f9e3
type TsFpUpdateBitmap struct {
	UpdateType       int16 // This field MUST be set to UPDATETYPE_BITMAP (0x0001).
	NumberRectangles uint16
	Rectangles       []TsBitmapData
}

func (t *TsFpUpdateBitmap) iUpdatePDU() {}

func (t *TsFpUpdateBitmap) Read(r io.Reader) UpdatePDU {
	core.ReadLE(r, &t.UpdateType)
	core.ReadLE(r, &t.NumberRectangles)
	glog.Debugf("number rectangles: %v", t.NumberRectangles)
	t.Rectangles = make([]TsBitmapData, t.NumberRectangles)
	for i := range t.Rectangles {
		glog.Debugf("number: %v", i)
		t.Rectangles[i].Read(r)
		glog.Debugf("rect: %v:%v - %v:%v (%x)",
			t.Rectangles[i].DestLeft, t.Rectangles[i].DestTop,
			t.Rectangles[i].DestRight, t.Rectangles[i].DestBottom, t.Rectangles[i].Flags)
	}
	glog.Debugf("UpdateBitMap read ok")
	return t
}
