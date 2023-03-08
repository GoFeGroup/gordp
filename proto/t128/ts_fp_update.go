package t128

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

// update code
const (
	FASTPATH_UPDATETYPE_ORDERS        = 0x0
	FASTPATH_UPDATETYPE_BITMAP        = 0x1
	FASTPATH_UPDATETYPE_PALETTE       = 0x2
	FASTPATH_UPDATETYPE_SYNCHRONIZE   = 0x3
	FASTPATH_UPDATETYPE_SURFCMDS      = 0x4
	FASTPATH_UPDATETYPE_PTR_NULL      = 0x5
	FASTPATH_UPDATETYPE_PTR_DEFAULT   = 0x6
	FASTPATH_UPDATETYPE_PTR_POSITION  = 0x8
	FASTPATH_UPDATETYPE_COLOR         = 0x9
	FASTPATH_UPDATETYPE_CACHED        = 0xA
	FASTPATH_UPDATETYPE_POINTER       = 0xB
	FASTPATH_UPDATETYPE_LARGE_POINTER = 0xC
)

// fragmentation
const (
	FASTPATH_FRAGMENT_SINGLE = 0x0
	FASTPATH_FRAGMENT_LAST   = 0x1
	FASTPATH_FRAGMENT_FIRST  = 0x2
	FASTPATH_FRAGMENT_NEXT   = 0x3
)

// compression
const (
	FASTPATH_OUTPUT_COMPRESSION_USED = 0x2
)

type UpdatePDU interface {
	iUpdatePDU()
	Read(r io.Reader) UpdatePDU
}

// FpOutputHeader
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/a1c4caa8-00ed-45bb-a06e-5177473766d3
type FpOutputHeader struct {
	UpdateCode    uint8
	Fragmentation uint8
	Compression   uint8
}

func (h *FpOutputHeader) Read(r io.Reader) {
	var updateHeader uint8
	core.ReadLE(r, &updateHeader)
	glog.Debugf("fpOutputHeader: %x", updateHeader)
	h.UpdateCode = updateHeader & 0xF
	h.Fragmentation = (updateHeader >> 4) & 0x03
	h.Compression = (updateHeader >> 6) & 0x03
	glog.Debugf("fpOutputHeader: %+v", h)
}

// TsFpUpdatePDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/68b5ee54-d0d5-4d65-8d81-e1c4025f7597
type TsFpUpdatePDU struct {
	Header FpOutputHeader
	Length uint16
	PDU    UpdatePDU
}

func (p *TsFpUpdatePDU) iPDU() {}

func (p *TsFpUpdatePDU) Serialize() []byte {
	//TODO implement me
	panic("implement me")
}

func (p *TsFpUpdatePDU) Type() uint16 {
	//TODO implement me
	panic("implement me")
}

func (p *TsFpUpdatePDU) Read(r io.Reader) PDU {
	p.Header.Read(r)
	if p.Header.Compression == FASTPATH_OUTPUT_COMPRESSION_USED {
		core.Throw("not implement")
	}
	core.ReadLE(r, &p.Length)
	if p.Length == 0 {
		glog.Debugf("length = 0")
		return p
	}

	data := core.ReadBytes(r, int(p.Length))
	//glog.Debugf("fastpath pdu data: %v - %x", len(data), data)

	glog.Debugf("updateCode: %v", p.Header.UpdateCode)
	switch p.Header.UpdateCode {
	case FASTPATH_UPDATETYPE_BITMAP:
		p.PDU = (&TsFpUpdateBitmap{}).Read(bytes.NewReader(data))
	default:
		glog.Warnf("updateCode [%x] not implement", p.Header.UpdateCode)
	}

	glog.Debugf("p.PDU: %T", p.PDU)
	return p
}
