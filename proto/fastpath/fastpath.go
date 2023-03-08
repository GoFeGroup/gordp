package fastpath

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/mcs/per"
	"io"
)

type Header struct {
	EncryptionFlags uint8
	NumberEvents    uint8
	Length          int
}

func (h *Header) Read(r io.Reader) {
	var b uint8
	core.ReadLE(r, &b)
	h.EncryptionFlags = (b & 0xc0) >> 6
	h.NumberEvents = (b & 0x3c) >> 2
	h.Length = per.ReadLength(r)
	h.Length = core.If(h.Length < 0x80, h.Length-2, h.Length-3)
}

type FastPathData struct {
	Header Header
	Data   []byte
}

func Read(r io.Reader) *FastPathData {
	fp := &FastPathData{}
	fp.Header.Read(r)
	//glog.Debugf("fastpath read header: %+v", fp.Header)
	fp.Data = core.ReadBytes(r, fp.Header.Length)
	//glog.Debugf("fastpath read data: %v - %x", len(fp.Data), fp.Data)
	return fp
}
