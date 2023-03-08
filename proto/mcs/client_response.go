package mcs

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs/ber"
	"io"
)

// ConnectResponse
// [T125] section 11.2 -- https://www.itu.int/rec/T-REC-T.125-199802-I/en
type ConnectResponse struct {
	Result           uint8
	CalledConnectId  int
	DomainParameters DomainParameters
	UserData         []byte
}

func (cr *ConnectResponse) Load(data []byte) {
	r := bytes.NewReader(data)
	userData := ber.ReadApplicationTag(r, MCS_TYPE_CONNECT_RESPONSE)

	glog.Debugf("ConnectResponse UserData: %v - %x", len(userData), userData)
	r = bytes.NewReader(userData)
	cr.Result = ber.ReadEnumerated(r)
	glog.Debugf("cr.result: %v", cr.Result)
	cr.CalledConnectId = ber.ReadInteger(r)
	glog.Debugf("cr.connectId: %v", cr.CalledConnectId)
	cr.DomainParameters.Read(r)

	core.ThrowIf(ber.ReadUniversalTag(r, ber.BER_TAG_OCTET_STRING, false) == false, "invalid universal tag")
	length := ber.ReadLength(r)
	cr.UserData = make([]byte, length)
	_, err := io.ReadFull(r, cr.UserData)
	core.ThrowError(err)
}
