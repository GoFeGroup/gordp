package t128

import (
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

// TsSaveSessionInfoPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/d892bc5b-aecd-4aee-99b6-5f43b5a63d75
type TsSaveSessionInfoPDU struct {
}

func (t *TsSaveSessionInfoPDU) iDataPDU() {}

func (t *TsSaveSessionInfoPDU) Read(r io.Reader) DataPDU {
	glog.Warnf("not implement")
	return t
}

func (t *TsSaveSessionInfoPDU) Serialize() []byte {
	//TODO implement me
	panic("implement me")
}

func (t *TsSaveSessionInfoPDU) Type2() uint8 {
	return PDUTYPE2_SAVE_SESSION_INFO
}
