package mcs

import (
	"github.com/GoFeGroup/gordp/proto/mcs/per"
	"io"
)

// for MCS Message
const (
	MCS_TYPE_CONNECT_INITIAL  = 0x65
	MCS_TYPE_CONNECT_RESPONSE = 0x66
)

// for MCS PDU type
const (
	MCS_PDUTYPE_ERECT_DOMAIN_REQUEST = 1
	MCS_PDUTYPE_ATTACH_USER_REQUEST  = 10
	MCS_PDUTYPE_ATTACH_USER_CONFIRM  = 11
	MCS_PDUTYPE_CHANNEL_JOIN_REQUEST = 14
	MCS_PDUTYPE_CHANNEL_JOIN_CONFIRM = 15
	MCS_PDUTYPE_SEND_DATA_REQUEST    = 25
	MCS_PDUTYPE_SEND_DATA_INDICATION = 26
)

const (
	MCS_CHANNEL_USERID_BASE = 1001
	MCS_CHANNEL_GLOBAL      = 1003
)

func WriteMcsPduHeader(w io.Writer, pduType, option uint8) {
	per.WriteChoice(w, (pduType<<2)|option)
}

func ReadMcsPduHeader(r io.Reader) uint8 {
	options := per.ReadChoice(r)
	return options >> 2
	//core.ThrowIf((options>>2) != pduType, "invalid pdu type")
}

//func Write(w io.Writer, channelId uint16, data []byte) {
//
//}
