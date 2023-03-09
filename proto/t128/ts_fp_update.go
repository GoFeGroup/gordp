package t128

import (
	"bytes"
	"io"

	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
)

type UpdatePDU interface {
	iUpdatePDU()
	Read(r io.Reader) UpdatePDU
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
