package mcs

import (
	"bytes"
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs/per"
	"github.com/GoFeGroup/gordp/proto/x224"
	"io"
)

type ReceiveDataResponse struct{}

func (res *ReceiveDataResponse) Read(r io.Reader) (uint16, []byte) {
	data := x224.Read(r)
	r = bytes.NewReader(data)
	pduHeader := ReadMcsPduHeader(r)
	core.ThrowIf(pduHeader != MCS_PDUTYPE_SEND_DATA_INDICATION, fmt.Errorf("invalid pdu header: %v", pduHeader))
	userId := per.ReadInteger16(r, MCS_CHANNEL_USERID_BASE) // UserId
	channelId := per.ReadInteger16(r, 0)
	glog.Debugf("userId: %v, channelId: %v", userId, channelId)
	enumerated := per.ReadEnumerated(r)
	glog.Debugf("enumerated: %v", enumerated)
	return channelId, per.ReadOctetString(r, 0)
}
