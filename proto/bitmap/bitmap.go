package bitmap

import (
	"bytes"
	"image"
	"image/png"

	"github.com/GoFeGroup/gordp/core"
)

import "C"

type Option struct {
	Top         int    `json:"top"`
	Left        int    `json:"left"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	BitPerPixel int    `json:"-"`
	Data        []byte `json:"-"`
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
