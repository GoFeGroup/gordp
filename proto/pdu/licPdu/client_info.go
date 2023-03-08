package licPdu

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/sec"
	"github.com/GoFeGroup/gordp/proto/x224"
	"io"
)

// ClientInfoPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/772d618e-b7d6-4cd0-b735-fa08af558f9d
type ClientInfoPDU struct {
	McsSDrq        *mcs.SendDataRequest // MCS Send Data Request
	SecurityHeader *sec.TsSecurityHeader
	InfoPacket     *TsInfoPacket
}

func NewClientInfoPDU(userId uint16, username, password string) *ClientInfoPDU {
	infoPkt := &TsInfoPacket{
		//Flag: INFO_MOUSE | INFO_UNICODE | INFO_LOGONERRORS | INFO_MAXIMIZESHELL |
		//	INFO_ENABLEWINDOWSKEY | INFO_DISABLECTRLALTDEL | INFO_MOUSE_HAS_WHEEL |
		//	INFO_FORCE_ENCRYPTED_CS_PDU | INFO_AUTOLOGON,
		Flag:           INFO_MOUSE | INFO_UNICODE | INFO_LOGONNOTIFY | INFO_LOGONERRORS | INFO_DISABLECTRLALTDEL | INFO_ENABLEWINDOWSKEY | INFO_AUTOLOGON,
		Domain:         []byte{0, 0},
		UserName:       append(core.UnicodeEncode(username), 0, 0),
		Password:       append(core.UnicodeEncode(password), 0, 0),
		AlternateShell: []byte{0, 0},
		WorkingDir:     []byte{0, 0},
		ExtendedInfo:   sec.NewExtendedInfoPacket(),
	}
	infoPkt.CbUserName = uint16(len(infoPkt.UserName) - 2)
	infoPkt.CbPassword = uint16(len(infoPkt.Password) - 2)

	return &ClientInfoPDU{
		McsSDrq:        mcs.NewSendDataRequest(userId, mcs.MCS_CHANNEL_GLOBAL),
		SecurityHeader: sec.NewTsSecurityHeader(sec.SEC_INFO_PKT),
		InfoPacket:     infoPkt,
	}
}

func (pdu *ClientInfoPDU) Serialize() []byte {
	buff := new(bytes.Buffer)
	pdu.SecurityHeader.Write(buff)
	pdu.InfoPacket.Write(buff)
	glog.Debugf("client info pdu data: %v - %x", buff.Len(), buff.Bytes())
	return buff.Bytes()
}

func (pdu *ClientInfoPDU) Write(w io.Writer) {
	data := pdu.McsSDrq.Serialize(pdu.Serialize())
	x224.Write(w, data)
}
