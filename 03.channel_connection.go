package gordp

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/pdu/mcsPdu"
)

func (c *Client) joinChannel(userId, channelId uint16) {
	mcsCJrq := mcsPdu.ClientMcsChannelJoinRequestPDU{}
	mcsCJrq.JoinChannel(c.stream, userId, channelId)

	mcsCJcf := mcsPdu.ServerMcsChannelJoinConfirmPDU{}
	mcsCJcf.Read(c.stream)

	core.ThrowIf(userId != mcsCJcf.McsCJcf.UserId, "invalid userId")
}

func (c *Client) channelConnect() {
	mcsEdrq := mcsPdu.ClientMcsErectDomainRequestPDU{}
	mcsEdrq.Write(c.stream)
	glog.Debugf("send erect domain request pdu ok")

	mcsAUrq := mcsPdu.ClientMcsAttachUserRequestPDU{}
	mcsAUrq.Write(c.stream)
	glog.Debugf("send attach user request pdu ok")

	mcsAUcf := mcsPdu.ServerMcsAttachUserConfirmPDU{}
	mcsAUcf.Read(c.stream)
	glog.Debugf("receive attach user confirm pdu ok")

	c.userId = mcsAUcf.McsAUcf.UserId

	c.joinChannel(c.userId, mcs.MCS_CHANNEL_GLOBAL) // join channel `global`
	c.joinChannel(c.userId, mcsAUcf.McsAUcf.UserId) // join channel `user`
}
