package t128

import (
	"bytes"

	"github.com/GoFeGroup/gordp/core"
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

// SlowPath Input PDU
// input_send_mouse_event
//  - rdp_client_input_pdu_init
//    - rdp_data_pdu_init
//      - rdp_send_stream_pdu_init
//        - rdp_send_stream_init
//          - RDP_PACKET_HEADER_MAX_LENGTH
//            = TPDU_DATA_LENGTH + MCS_SEND_DATA_HEADER_MAX_LENGTH
//            = (TPKT_HEADER_LENGTH + TPDU_DATA_HEADER_LENGTH) + 8
//            = (4 + 3) + 8
//          - security == can be 0
//        - RDP_SHARE_CONTROL_HEADER_LENGTH = 6
//     - RDP_SHARE_DATA_HEADER_LENGTH = 12
//    - rdp_write_client_input_pdu_header   // TS_INPUT_PDU_DATA <- SlowPath
//      - numberEvents = 2
//      - pad2Octets = 2
//    - rdp_write_input_event_header
//      - eventTime = 4
//      - messageType = 2
//  - input_write_mouse_event
//    - flags = 2
//    - xPos = 2
//    - yPos = 2

// FastPath Input PDU
// input_send_fastpath_mouse_event
//  - fastpath_input_pdu_init
//    - fastpath_input_pdu_init_header
//      - transport_send_stream_init = 0
//      - fpInputHeader, length1 and length2 = 3
//      - fastpath_get_sec_bytes = 0
//    - eventHeader = (eventFlags | (eventCode << 5)) = 1
//  - input_write_mouse_event
//    - flags = 2
//    - xPos = 2
//    - yPos = 2

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

	core.WriteBE(buff, (pdu.Length+3)|0x8000) // copy from FreeRDP
	//per.WriteLength(buff, int(pdu.Length))
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
