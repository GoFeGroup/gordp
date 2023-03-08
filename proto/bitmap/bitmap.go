package bitmap

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"image"
	"image/png"
)

import "C"

type Option struct {
	Width       int
	Height      int
	BitPerPixel int
	Data        []byte
}

type BitMap struct {
	Image image.Image
}

func (m *BitMap) ToPng() []byte {
	buf := new(bytes.Buffer)
	core.ThrowError(png.Encode(buf, m.Image))
	return buf.Bytes()
}

func NewBitMapFromRDP6(option *Option) *BitMap {
	return (&BitMap{}).LoadRDP60(option)
}

func NewBitmapFromRLE(option *Option) *BitMap {
	return (&BitMap{}).LoadRLE(option)
}
