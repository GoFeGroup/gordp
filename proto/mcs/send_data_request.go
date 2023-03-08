package mcs

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/mcs/per"
	"io"
)

// SendDataRequest --- Copy From FreeRDP C.rdp_write_header
type SendDataRequest struct {
	UserId    uint16
	ChannelId uint16
}

func (r *SendDataRequest) Write(w io.Writer, data []byte) {
	WriteMcsPduHeader(w, MCS_PDUTYPE_SEND_DATA_REQUEST, 0)
	core.WriteBE(w, r)
	core.WriteBE(w, uint8(0x70)) // dataPriority + segmentation
	per.WriteLength(w, len(data))
	core.WriteFull(w, data)
}

func (r *SendDataRequest) Serialize(data []byte) []byte {
	buff := new(bytes.Buffer)
	r.Write(buff, data)
	return buff.Bytes()
}

func NewSendDataRequest(userId, channelId uint16) *SendDataRequest {
	return &SendDataRequest{
		UserId:    userId - MCS_CHANNEL_USERID_BASE,
		ChannelId: channelId,
	}
}
