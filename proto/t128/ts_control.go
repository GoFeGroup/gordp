package t128

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// Action
const (
	CTRLACTION_REQUEST_CONTROL = 0x0001 //Request control
	CTRLACTION_GRANTED_CONTROL = 0x0002 //Granted control
	CTRLACTION_DETACH          = 0x0003 //Detach
	CTRLACTION_COOPERATE       = 0x0004 //Cooperate
)

// TsControlPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/0448f397-aa11-455d-81b1-f1265085239d
type TsControlPDU struct {
	Action    uint16
	GrantId   uint16
	ControlId uint32
}

func (t *TsControlPDU) iDataPDU() {}

func (t *TsControlPDU) Read(r io.Reader) DataPDU {
	return core.ReadLE(r, t)
}

func (t *TsControlPDU) Serialize() []byte {
	return core.ToLE(t)
}

func (t *TsControlPDU) Type2() uint8 {
	return PDUTYPE2_CONTROL
}
