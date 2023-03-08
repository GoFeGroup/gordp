package nla

import (
	"bytes"
	"crypto/md5"
	"crypto/rc4"
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// AuthenticateMessage 认证信息
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/033d32cc-88f9-4483-9bf2-b273055038ce?source=recommendations
type AuthenticateMessage struct {
	Must struct {
		Signature   [8]byte // MUST contain the ASCII string ('N', 'T', 'L', 'M', 'S', 'S', 'P', '\0').
		MessageType uint32  //  This field MUST be set to 0x00000003.

		LmChallengeResponse    Field
		NtChallengeResponse    Field
		DomainName             Field
		UserName               Field
		Workstation            Field
		EncryptedRandomSession Field

		NegotiateFlags uint32
		Version        NVersion // 8bytes
		MIC            [16]byte
	}

	Optional struct {
		Payload []byte // variable

		user, pass string
		offset     uint32

		NtlmSec       *NTLMv2Security
		encryptPubkey []byte
	}
}

func (m *AuthenticateMessage) Serialize() []byte {
	buff := new(bytes.Buffer)
	core.WriteLE(buff, &m.Must)
	buff.Write(m.Optional.Payload)
	return buff.Bytes()
}

func (m *AuthenticateMessage) BaseLen() uint32 {
	return 88
}

func (m *AuthenticateMessage) SetNegotiateFlags(flags uint32) *AuthenticateMessage {
	m.Must.NegotiateFlags = flags
	return m
}

func (m *AuthenticateMessage) SetLmChallengeResponse(length uint16) *AuthenticateMessage {
	m.Must.LmChallengeResponse.Set(length, m.Optional.offset)
	m.Optional.offset += uint32(length)
	return m
}

func (m *AuthenticateMessage) SetNtChallengeResponse(length uint16) *AuthenticateMessage {
	m.Must.NtChallengeResponse.Set(length, m.Optional.offset)
	m.Optional.offset += uint32(length)
	return m
}

func (m *AuthenticateMessage) SetDomainName(length uint16) *AuthenticateMessage {
	m.Must.DomainName.Set(length, m.Optional.offset)
	m.Optional.offset += uint32(length)
	return m
}

func (m *AuthenticateMessage) SetUserName(length uint16) *AuthenticateMessage {
	m.Must.UserName.Set(length, m.Optional.offset)
	m.Optional.offset += uint32(length)
	return m
}

func (m *AuthenticateMessage) SetWorkstation(length uint16) *AuthenticateMessage {
	m.Must.Workstation.Set(length, m.Optional.offset)
	m.Optional.offset += uint32(length)
	return m
}

func (m *AuthenticateMessage) SetEncryptedRandomSession(length uint16) *AuthenticateMessage {
	m.Must.EncryptedRandomSession.Set(length, m.Optional.offset)
	m.Optional.offset += uint32(length)
	return m
}

func MIC(exportedSessionKey []byte, negotiate *NegotiateMessage, challenge *ChallengeMessage, auth *AuthenticateMessage) []byte {
	buff := bytes.Buffer{}
	buff.Write(negotiate.Serialize())
	buff.Write(challenge.Serialize())
	buff.Write(auth.Serialize())
	return core.HMAC_MD5(exportedSessionKey, buff.Bytes())
}

var (
	clientSigning = concat([]byte("session key to client-to-server signing key magic constant"), []byte{0x00})
	serverSigning = concat([]byte("session key to server-to-client signing key magic constant"), []byte{0x00})
	clientSealing = concat([]byte("session key to client-to-server sealing key magic constant"), []byte{0x00})
	serverSealing = concat([]byte("session key to server-to-client sealing key magic constant"), []byte{0x00})
)

// CalcChallenge
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/c0250a97-2940-40c7-82fb-20d208c71e96
func (m *AuthenticateMessage) CalcChallenge(negotiate *NegotiateMessage, challenge *ChallengeMessage) *AuthenticateMessage {
	respKeyNT := core.NTOWFv2(m.Optional.pass, m.Optional.user, "")
	respKeyLM := core.LMOWFv2(m.Optional.pass, m.Optional.user, "")

	ntChallenge := NewNTLMv2ClientChallenge(challenge.getTargetInfo())
	ccData := ntChallenge.Serialize()

	serverChallenge := challenge.Must.ServerChallenge[:]
	clientChallenge := ntChallenge.Must.ChallengeFromClient[:]

	ntProof := core.HMAC_MD5(respKeyNT, append(serverChallenge, ccData...))
	ntChallResp := append(ntProof, ccData...)

	lmProof := core.HMAC_MD5(respKeyLM, append(serverChallenge, clientChallenge...))
	lmChallResp := append(lmProof, clientChallenge...)

	sessBasekey := core.HMAC_MD5(respKeyNT, ntProof)

	exportedSessionKey := core.Random(16)
	EncryptedRandomSessionKey := make([]byte, len(exportedSessionKey))
	rc, _ := rc4.NewCipher(sessBasekey)
	rc.XORKeyStream(EncryptedRandomSessionKey, exportedSessionKey)

	var user = []byte(m.Optional.user)
	if challenge.Must.NegotiateFlags&NTLMSSP_NEGOTIATE_UNICODE != 0 {
		user = core.UnicodeEncode(m.Optional.user)
	}

	m.SetNegotiateFlags(challenge.Must.NegotiateFlags).
		SetLmChallengeResponse(uint16(len(lmChallResp))).
		SetNtChallengeResponse(uint16(len(ntChallResp))).
		SetDomainName(uint16(len(""))).
		SetUserName(uint16(len(user))).
		SetWorkstation(uint16(len(""))).
		SetEncryptedRandomSession(uint16(len(EncryptedRandomSessionKey)))

	buff := new(bytes.Buffer)
	buff.Write(lmChallResp)
	buff.Write(ntChallResp)
	buff.Write([]byte("")) // domain
	buff.Write(user)
	buff.Write([]byte("")) // workstation
	buff.Write(EncryptedRandomSessionKey)

	if (m.Must.NegotiateFlags & NTLMSSP_NEGOTIATE_VERSION) != 0 {
		m.Must.Version = NewNVersion()
	}
	m.Optional.Payload = buff.Bytes()

	// calculate MIC
	copy(m.Must.MIC[:], MIC(exportedSessionKey, negotiate, challenge, m))

	ClientSigningKey := md5.Sum(concat(exportedSessionKey, clientSigning))
	ServerSigningKey := md5.Sum(concat(exportedSessionKey, serverSigning))
	ClientSealingKey := md5.Sum(concat(exportedSessionKey, clientSealing))
	ServerSealingKey := md5.Sum(concat(exportedSessionKey, serverSealing))

	encryptRC4, _ := rc4.NewCipher(ClientSealingKey[:])
	decryptRC4, _ := rc4.NewCipher(ServerSealingKey[:])

	m.Optional.NtlmSec = &NTLMv2Security{encryptRC4, decryptRC4, ClientSigningKey[:], ServerSigningKey[:], 0}
	return m
}

func (m *AuthenticateMessage) Sign(pubKey []byte) *AuthenticateMessage {
	m.Optional.encryptPubkey = m.Optional.NtlmSec.Serialize(pubKey)
	return m
}

func (m *AuthenticateMessage) Write(w io.Writer) {
	req := NewTsRequest().SetMessages(m.Serialize()).SetPubKeyAuth(m.Optional.encryptPubkey)
	req.Write(w)
}

func NewAuthenticateMessage(user, pass string) *AuthenticateMessage {
	m := &AuthenticateMessage{}
	m.Must.Signature = [8]byte{'N', 'T', 'L', 'M', 'S', 'S', 'P', 0x00}
	m.Must.MessageType = 0x00000003
	m.Optional.offset = m.BaseLen()
	m.Optional.user = user
	m.Optional.pass = pass
	return m
}

//
//func NewAuthenticateMessage(negFlag uint32, domain, user, workstation []byte,
//	lmchallResp, ntchallResp, enRandomSessKey []byte) *AuthenticateMessage {
//	msg := &AuthenticateMessage{}
//	msg.Must.Signature = [8]byte{'N', 'T', 'L', 'M', 'S', 'S', 'P', 0x00}
//	msg.Must.MessageType = 0x00000003
//	msg.Optional.offset = msg.BaseLen()
//
//	msg = msg.SetNegotiateFlags(negFlag).SetLmChallengeResponse(uint16(len(lmchallResp))).
//		SetNtChallengeResponse(uint16(len(ntchallResp))).
//		SetDomainName(uint16(len(domain))).SetUserName(uint16(len(user))).
//		SetWorkstation(uint16(len(workstation))).SetEncryptedRandomSession(uint16(len(enRandomSessKey)))
//
//	buff := new(bytes.Buffer)
//	buff.Write(lmchallResp)
//	buff.Write(ntchallResp)
//	buff.Write(domain)
//	buff.Write(user)
//	buff.Write(workstation)
//	buff.Write(enRandomSessKey)
//
//	if (msg.Must.NegotiateFlags & NTLMSSP_NEGOTIATE_VERSION) != 0 {
//		msg.Must.Version = NewNVersion()
//	}
//	msg.Optional.Payload = buff.Bytes()
//	return msg
//}
