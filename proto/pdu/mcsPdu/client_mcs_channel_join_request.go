package mcsPdu

import (
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/x224"
	"io"
)

// ClientMcsChannelJoinRequestPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/64564639-3b2d-4d2c-ae77-1105b4cc011b
type ClientMcsChannelJoinRequestPDU struct {
}

func (pdu *ClientMcsChannelJoinRequestPDU) JoinChannel(w io.Writer, userId, channelId uint16) {
	mcsCJrq := mcs.NewClientChannelJoin(userId, channelId)
	data := mcsCJrq.Serialize()
	glog.Debugf("send join channel request: %v - %x", len(data), data)
	x224.Write(w, data)
}
