package t128

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsSetErrorInfoPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/a21a1bd9-2303-49c1-90ec-3932435c248c
type TsSetErrorInfoPDU struct {
	ErrorInfo uint32
}

func (t *TsSetErrorInfoPDU) iDataPDU() {}

func (t *TsSetErrorInfoPDU) Read(r io.Reader) DataPDU {
	return core.ReadLE(r, t)
}

func (t *TsSetErrorInfoPDU) Serialize() []byte {
	//TODO implement me
	panic("implement me")
}

func (t *TsSetErrorInfoPDU) Type2() uint8 {
	return PDUTYPE2_SET_ERROR_INFO_PDU
}
