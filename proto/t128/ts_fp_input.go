package t128

import (
	"bytes"

	"github.com/GoFeGroup/gordp/proto/mcs/per"
)

// TsFpInputPdu
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/b8e7c588-51cb-455b-bb73-92d480903133
type TsFpInputPdu struct {
	Header          FpInputHeader
	Length          uint16
	FipsInformation uint32           // Optional: when Server Security Data (TS_UD_SC_SEC1) is set
	DataSignature   [8]byte          // Optional: existed if (Header.Flag & FASTPATH_INPUT_SECURE_CHECKSUM)
	NumEvents       uint8            // Optional: if (header.NumEvent != 0)
	FpInputEvents   []TsFpInputEvent // An array of Fast-Path Input Event (section 2.2.8.1.2.2)
}

//func (pdu *TsFpInputPdu) iDataPDU() {}
//
//func (pdu *TsFpInputPdu) Read(r io.Reader) DataPDU {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (pdu *TsFpInputPdu) Type2() uint8 {
//	return PDUTYPE2_INPUT
//}

func (pdu *TsFpInputPdu) Serialize() []byte {
	var events [][]byte
	for _, v := range pdu.FpInputEvents {
		events = append(events, v.Serialize())
	}
	eventsData := bytes.Join(events, nil)
	pdu.Length = uint16(len(eventsData))

	pdu.Header.Action = FASTPATH_INPUT_ACTION_FASTPATH
	pdu.Header.NumEvents = uint8(len(pdu.FpInputEvents))

	buff := new(bytes.Buffer)
	pdu.Header.Write(buff)

	per.WriteLength(buff, int(pdu.Length))
	buff.Write(eventsData)

	return buff.Bytes()
}

func NewFastPathMouseInputPDU(pointerFlags uint16, xPos, yPos uint16) *TsFpInputPdu {
	return &TsFpInputPdu{
		FpInputEvents: []TsFpInputEvent{&TsFpPointerEvent{
			PointerFlags: pointerFlags,
			XPos:         xPos,
			YPos:         yPos,
		}},
	}
}
