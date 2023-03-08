package t128

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

type TsDataPduData struct {
	Header  TsShareDataHeader
	PduData []byte
	Pdu     DataPDU
}

func (t *TsDataPduData) iPDU() {}

func (t *TsDataPduData) Type() uint16 {
	return PDUTYPE_DATAPDU
}

func (t *TsDataPduData) Read(r io.Reader) PDU {
	t.Header.Read(r)
	glog.Debugf("data header: %+v", t.Header)
	t.Pdu = pduMap2[t.Header.PDUType2].Read(r)
	return t
}

func (t *TsDataPduData) Serialize() []byte {
	buff := new(bytes.Buffer)
	core.WriteLE(buff, t.Header)
	core.WriteFull(buff, t.PduData)
	return buff.Bytes()
}

func NewDataPdu(pdu DataPDU, shareId uint32) *TsDataPduData {
	dataPdu := &TsDataPduData{}
	dataPdu.PduData = pdu.Serialize()
	glog.Debugf("data@DataPDU: %v - %x", len(dataPdu.PduData), dataPdu.PduData)

	dataPdu.Header.StreamId = STREAM_LOW
	dataPdu.Header.SharedId = shareId
	dataPdu.Header.PDUType2 = pdu.Type2()
	dataPdu.Header.UncompressedLength = uint16(len(dataPdu.PduData) + 4)
	return dataPdu
}
