package t128

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

const (
	SYNCMSGTYPE_SYNC = 1
)

// TsSynchronizePduData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/3fb4c95e-ad2d-43d1-a46f-5bd49418da49
type TsSynchronizePduData struct {
	MessageType uint16 // This field MUST be set to SYNCMSGTYPE_SYNC (1).
	TargetUser  uint16 //A 16-bit, unsigned integer. The MCS channel ID of the target user.
}

func (t *TsSynchronizePduData) Read(r io.Reader) DataPDU {
	return core.ReadLE(r, t)
}

func (t *TsSynchronizePduData) Type2() uint8 {
	return PDUTYPE2_SYNCHRONIZE
}

func (t *TsSynchronizePduData) iDataPDU() {}

func (t *TsSynchronizePduData) Serialize() []byte {
	return core.ToLE(t)
}

func NewTsSynchronizePduData(targetUser uint16) *TsSynchronizePduData {
	return &TsSynchronizePduData{
		MessageType: SYNCMSGTYPE_SYNC,
		TargetUser:  targetUser,
	}
}
