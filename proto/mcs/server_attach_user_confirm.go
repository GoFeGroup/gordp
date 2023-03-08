package mcs

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs/per"
	"io"
)

type ServerAttachUserConfirm struct {
	UserId uint16
}

func (c *ServerAttachUserConfirm) Read(r io.Reader) {
	core.ThrowIf(ReadMcsPduHeader(r) != MCS_PDUTYPE_ATTACH_USER_CONFIRM, "invalid pdu TYPE")
	core.ThrowIf(per.ReadEnumerated(r) != 0, "invalid enumerated")
	c.UserId = per.ReadInteger16(r, 0) + MCS_CHANNEL_USERID_BASE // userId base
	glog.Debugf("userId: %v", c.UserId)
}
