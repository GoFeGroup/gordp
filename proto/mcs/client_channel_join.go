package mcs

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// ClientChannelJoin
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/64564639-3b2d-4d2c-ae77-1105b4cc011b
type ClientChannelJoin struct {
	UserId    uint16
	ChannelId uint16
}

func NewClientChannelJoin(userId, channelId uint16) *ClientChannelJoin {
	return &ClientChannelJoin{
		UserId:    userId - MCS_CHANNEL_USERID_BASE,
		ChannelId: channelId,
	}
}
func (j *ClientChannelJoin) Write(w io.Writer) {
	WriteMcsPduHeader(w, MCS_PDUTYPE_CHANNEL_JOIN_REQUEST, 0)
	core.WriteBE(w, j)
}

func (j *ClientChannelJoin) Serialize() []byte {
	buff := new(bytes.Buffer)
	j.Write(buff)
	return buff.Bytes()
}
