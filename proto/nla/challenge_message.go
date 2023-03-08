package nla

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
	"time"
)

type ChallengeMessage struct {
	msgBase
	Must struct {
		Signature       [8]byte // MUST contain the ASCII string ('N', 'T', 'L', 'M', 'S', 'S', 'P', '\0').
		MessageType     uint32  // This field MUST be set to 0x00000002.
		TargetName      Field   // NTLMSSP_REQUEST_TARGET
		NegotiateFlags  uint32
		ServerChallenge [8]byte //  A 64-bit value that contains the NTLM challenge. The challenge is a 64-bit nonce.
		Reserved        [8]byte
		TargetInfo      Field // NTLMSSP_NEGOTIATE_TARGET_INFO
	}
	Optional struct {
		Version NVersion // NTLMSSP_NEGOTIATE_VERSION
		Payload []byte
	}
}

func (m *ChallengeMessage) BaseLen() uint32 {
	if m.Must.NegotiateFlags&NTLMSSP_NEGOTIATE_VERSION != 0 {
		return 56 // include Version
	}
	return 48 // exclude Version
}

func (m *ChallengeMessage) Serialize() []byte {
	buff := new(bytes.Buffer)
	core.WriteLE(buff, &m.Must)
	if (m.Must.NegotiateFlags & NTLMSSP_NEGOTIATE_VERSION) != 0 {
		core.WriteLE(buff, &m.Optional.Version)
	}
	buff.Write(m.Optional.Payload)
	return buff.Bytes()
}

func (m *ChallengeMessage) Load(r *bytes.Reader) {
	core.ReadLE(r, &m.Must)
	core.ThrowIf(m.Must.MessageType != 0x00000002, fmt.Errorf("invalid message type: %x", m.Must.MessageType))
	if m.Must.NegotiateFlags&NTLMSSP_NEGOTIATE_VERSION != 0 {
		core.ReadLE(r, &m.Optional.Version)
	}

	m.Optional.Payload = core.ReadBytes(r, r.Len())
}

func (m *ChallengeMessage) Read(r io.Reader) {
	req := &TSRequest{}
	req.Read(r)
	core.ThrowIf(len(req.NegoTokens) == 0, fmt.Errorf("invalid TsRequest from upstream, NegoTokens = nil"))
	m.Load(bytes.NewReader(req.NegoTokens[0].Data))
}

//func (m *ChallengeMessage) Read(r *bytes.Reader) error {
//	if err := binary.Read(r, binary.LittleEndian, &m.Must); err != nil {
//		return err
//	}
//	if m.Must.NegotiateFlags&NTLMSSP_NEGOTIATE_VERSION != 0 {
//		if err := binary.Read(r, binary.LittleEndian, &m.Optional.Version); err != nil {
//			return err
//		}
//	}
//
//	m.Optional.Payload = make([]byte, r.Len())
//	_, err := io.ReadFull(r, m.Optional.Payload)
//	return err
//}

func (m *ChallengeMessage) getTargetName() []byte {
	return m.GetField(m.Optional.Payload, m.BaseLen(), &m.Must.TargetName)
}

func (m *ChallengeMessage) getTimestamp(data []byte) []byte {
	avPairs := ReadAvPairs(data)
	return avPairs.GetTimeStamp()
}

func (m *ChallengeMessage) getTargetInfo() ([]byte, []byte) {
	data := m.GetField(m.Optional.Payload, m.BaseLen(), &m.Must.TargetInfo)
	glog.Debugf("targetInfo: %x", data)
	if tm := m.getTimestamp(data); tm != nil {
		glog.Debugf("get timestamp %v", tm)
		return data, tm
	}
	ft := uint64(time.Now().UnixNano()) / 100
	ft += 116444736000000000 // add time between unix & windows offset
	tm := make([]byte, 8)
	binary.LittleEndian.PutUint64(tm, ft)
	return data, tm
}
