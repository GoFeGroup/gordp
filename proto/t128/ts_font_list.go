package t128

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsFontListPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/e373575a-01e2-43a7-a6d8-e1952b83e787
type TsFontListPDU struct {
	NumberFonts   uint16
	TotalNumFonts uint16
	ListFlags     uint16 //This field SHOULD be set to 0x0003
	EntrySize     uint16 //This field SHOULD be set to 0x0032 (50 bytes).
}

func (t *TsFontListPDU) Read(r io.Reader) DataPDU {
	//TODO implement me
	panic("implement me")
}

func (t *TsFontListPDU) iDataPDU() {}

func (t *TsFontListPDU) Serialize() []byte {
	return core.ToLE(t)
}

func (t *TsFontListPDU) Type2() uint8 {
	return PDUTYPE2_FONTLIST
}
