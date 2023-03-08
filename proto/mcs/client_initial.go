package mcs

import (
	"bytes"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs/ber"
)

// ConnectInitial
// [T125] section 11.1 -- https://www.itu.int/rec/T-REC-T.125-199802-I/en
type ConnectInitial struct {
	CallingDomainSelector   []byte // 0x01
	CalledDomainSelector    []byte // 0x01
	UpwardFlag              bool   // true
	TargetDomainParameters  DomainParameters
	MinimumDomainParameters DomainParameters
	MaximumDomainParameters DomainParameters
	UserData                []byte
}

func (ci *ConnectInitial) Write(w *bytes.Buffer) {
	ber.WriteOctetstring(w, string(ci.CallingDomainSelector))
	ber.WriteOctetstring(w, string(ci.CalledDomainSelector))
	ber.WriteBoolean(w, ci.UpwardFlag)
	ber.WriteDomainParameters(w, ci.TargetDomainParameters.Serialize())
	ber.WriteDomainParameters(w, ci.MinimumDomainParameters.Serialize())
	ber.WriteDomainParameters(w, ci.MaximumDomainParameters.Serialize())
	ber.WriteOctetstring(w, string(ci.UserData))
}

func (ci *ConnectInitial) Serialize() []byte {
	buff := &bytes.Buffer{}
	ci.Write(buff)
	glog.Debugf("ConnectInitial: %v", buff.Len())

	buff2 := new(bytes.Buffer)
	ber.WriteApplicationTag(buff2, uint8(MCS_TYPE_CONNECT_INITIAL), buff.Bytes())
	glog.Debugf("x224 write: %v", buff2.Len())
	return buff2.Bytes()
}

func NewClientInitial() *ConnectInitial {
	return &ConnectInitial{
		CallingDomainSelector:   []byte{0x1},
		CalledDomainSelector:    []byte{0x1},
		UpwardFlag:              true,
		TargetDomainParameters:  DomainParameters{34, 2, 0, 1, 0, 1, 0xffff, 2},
		MinimumDomainParameters: DomainParameters{1, 1, 1, 1, 0, 1, 0x420, 2},
		MaximumDomainParameters: DomainParameters{0xffff, 0xfc17, 0xffff, 1, 0, 1, 0xffff, 2},
	}
}
