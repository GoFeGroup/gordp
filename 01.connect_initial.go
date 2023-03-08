package gordp

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/nla"
	"github.com/GoFeGroup/gordp/proto/pdu/connPdu"
)

func (c *Client) switchNLA() {
	// 发送 NegotiateMessage
	negotiate := nla.NewNegotiateMessage()
	negotiate.Write(c.stream)
	glog.Debugf("send negotiate message ok.")

	// 读取 ChallengeMessage
	challenge := &nla.ChallengeMessage{}
	challenge.Read(c.stream)
	glog.Debugf("recv challenge message ok")

	// 发送 AuthenticateMessage
	pk := c.stream.PubKey()
	auth := nla.NewAuthenticateMessage(c.option.UserName, c.option.Password)
	auth.CalcChallenge(negotiate, challenge).Sign(pk).Write(c.stream)

	// 读取 PubKeyAuth
	tsReq := &nla.TSRequest{}
	tsReq.Read(c.stream)
	glog.Debug("PubKeyAuth:", tsReq.PubKeyAuth)

	// 发送 Credentials
	tpCred := nla.TSPasswordCreds{
		DomainName: []byte(""),
		UserName:   []byte(c.option.UserName),
		Password:   []byte(c.option.Password),
	}
	if challenge.Must.NegotiateFlags&nla.NTLMSSP_NEGOTIATE_UNICODE != 0 {
		tpCred.UserName = core.UnicodeEncode(c.option.UserName)
		tpCred.Password = core.UnicodeEncode(c.option.Password)
	}
	tCred := nla.TSCredentials{CredType: 1, Credentials: tpCred.Serialize()}
	authInfo := auth.Optional.NtlmSec.Serialize(tCred.Serialize())
	nla.NewTsRequest().SetAuthInfo(authInfo).Write(c.stream)
}

// Connection Sequence
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/023f1e69-cfe8-4ee6-9ee0-7e759fb4e4ee
func (c *Client) negotiation() {
	reqPdu := connPdu.NewClientConnectionRequestPDU()
	reqPdu.Write(c.stream)

	resPdu := &connPdu.ServerConnectionConfirmPDU{}
	resPdu.Read(c.stream)

	switch resPdu.ProtocolNeg.Result {
	case connPdu.PROTOCOL_RDP:
	case connPdu.PROTOCOL_SSL:
		c.stream.SwitchSSL()
	case connPdu.PROTOCOL_HYBRID:
		c.stream.SwitchSSL()
		c.switchNLA()
	default:
		core.Throw("invalid protocol")
	}

	c.selectProtocol = resPdu.ProtocolNeg.Result
}
