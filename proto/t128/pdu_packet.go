package t128

import (
	"bytes"
	"fmt"
	"io"

	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/fastpath"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/x224"
)

type PDU interface {
	iPDU()
	Read(r io.Reader) PDU
	Serialize() []byte
	Type() uint16
}

type DataPDU interface {
	iDataPDU()
	Read(r io.Reader) DataPDU
	Serialize() []byte
	Type2() uint8
}

var pduMap = map[uint16]PDU{
	PDUTYPE_DEMANDACTIVEPDU:  &TsDemandActivePduData{},
	PDUTYPE_CONFIRMACTIVEPDU: &TsConfirmActivePduData{},
	PDUTYPE_DEACTIVATEALLPDU: nil,
	PDUTYPE_DATAPDU:          &TsDataPduData{},
	PDUTYPE_SERVER_REDIR_PKT: nil,
}

var pduMap2 = map[uint8]DataPDU{
	PDUTYPE2_SYNCHRONIZE:        &TsSynchronizePduData{},
	PDUTYPE2_CONTROL:            &TsControlPDU{},
	PDUTYPE2_FONTMAP:            &TsFontMapPDU{},
	PDUTYPE2_SET_ERROR_INFO_PDU: &TsSetErrorInfoPDU{},
	PDUTYPE2_SAVE_SESSION_INFO:  &TsSaveSessionInfoPDU{},
}

func readPDU(r io.Reader, typ uint16) PDU {
	if _, ok := pduMap[typ]; !ok {
		core.Throw(fmt.Errorf("invalid pdu type: %v", typ))
	}
	return pduMap[typ].Read(r)
}

func readMcsSdin(r io.Reader) []byte {
	var mcsSDin mcs.ReceiveDataResponse
	channelId, data := mcsSDin.Read(r)
	glog.Debugf("read pdu from channel: %v, %x", channelId, data)
	return data
}

func ReadExpectedPDU(r io.Reader, typ uint16) PDU {
	r = bytes.NewReader(readMcsSdin(r))
	header := TsShareControlHeader{}
	header.Read(r)
	glog.Debugf("share ctrl header: %+v", header)
	core.ThrowIf(header.PDUType != typ, "not expected PDU type")
	return readPDU(r, typ)
}

func ReadPDU(r io.Reader) PDU {
	r = bytes.NewReader(readMcsSdin(r))
	header := TsShareControlHeader{}
	header.Read(r)
	return readPDU(r, header.PDUType)
}

func WritePDU(w io.Writer, userId uint16, pdu PDU) {
	data := pdu.Serialize()
	header := TsShareControlHeader{
		PDUType:     pdu.Type(),
		PDUSource:   userId,
		TotalLength: uint16(len(data) + 6),
	}
	glog.Debugf("pdu.Serialize: %v - %x", len(data), data)

	mcsSDrq := mcs.NewSendDataRequest(userId, mcs.MCS_CHANNEL_GLOBAL)
	data = mcsSDrq.Serialize(append(header.Serialize(), data...))
	x224.Write(w, data)
}

func ReadExpectedDataPDU(r io.Reader, typ2 uint8) DataPDU {
	pdu := ReadExpectedPDU(r, PDUTYPE_DATAPDU).(*TsDataPduData)
	core.ThrowIf(pdu.Header.PDUType2 != typ2, "invalid pdu type2")
	return pdu.Pdu
}

func WriteDataPdu(w io.Writer, userId uint16, shareId uint32, pdu DataPDU) {
	WritePDU(w, userId, NewDataPdu(pdu, shareId))
}

func ReadFastPathPDU(r io.Reader) PDU {
	fp := fastpath.Read(r)
	if fp.Header.EncryptionFlags != 0 {
		core.Throw("not implement")
	}
	glog.Debugf("analyse FastPathPDU")
	return (&TsFpUpdatePDU{}).Read(bytes.NewReader(fp.Data))
}

func WriteFastPathPDU(w io.Writer, pdu PDU) {
	data := pdu.Serialize()
	fastpath.Write(w, data)
}
