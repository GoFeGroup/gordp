package mcs

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/mcs/per"
	"io"
)

type ServerChannelJoinConfirm struct {
	Confirm   uint8
	UserId    uint16
	ChannelId uint16
}

func (c *ServerChannelJoinConfirm) Read(r io.Reader) {
	core.ThrowIf(ReadMcsPduHeader(r) != MCS_PDUTYPE_CHANNEL_JOIN_CONFIRM, "invalid pdu Type")
	c.Confirm = per.ReadEnumerated(r)
	c.UserId = per.ReadInteger16(r, 0) + MCS_CHANNEL_USERID_BASE
	c.ChannelId = per.ReadInteger16(r, 0)

	core.ThrowIf(c.Confirm != 0 && (c.ChannelId == MCS_CHANNEL_GLOBAL || c.ChannelId == c.UserId), "not confirm")
}
