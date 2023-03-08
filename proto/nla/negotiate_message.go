package nla

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/99d90ff4-957f-4c8a-80e4-5bfe5a9a9832
const (
	NTLMSSP_NEGOTIATE_56                       = 0x80000000
	NTLMSSP_NEGOTIATE_KEY_EXCH                 = 0x40000000
	NTLMSSP_NEGOTIATE_128                      = 0x20000000
	NTLMSSP_NEGOTIATE_VERSION                  = 0x02000000
	NTLMSSP_NEGOTIATE_TARGET_INFO              = 0x00800000
	NTLMSSP_REQUEST_NON_NT_SESSION_KEY         = 0x00400000
	NTLMSSP_NEGOTIATE_IDENTIFY                 = 0x00100000
	NTLMSSP_NEGOTIATE_EXTENDED_SESSIONSECURITY = 0x00080000
	NTLMSSP_TARGET_TYPE_SERVER                 = 0x00020000
	NTLMSSP_TARGET_TYPE_DOMAIN                 = 0x00010000
	NTLMSSP_NEGOTIATE_ALWAYS_SIGN              = 0x00008000
	NTLMSSP_NEGOTIATE_OEM_WORKSTATION_SUPPLIED = 0x00002000
	NTLMSSP_NEGOTIATE_OEM_DOMAIN_SUPPLIED      = 0x00001000
	NTLMSSP_NEGOTIATE_NTLM                     = 0x00000200
	NTLMSSP_NEGOTIATE_LM_KEY                   = 0x00000080
	NTLMSSP_NEGOTIATE_DATAGRAM                 = 0x00000040
	NTLMSSP_NEGOTIATE_SEAL                     = 0x00000020
	NTLMSSP_NEGOTIATE_SIGN                     = 0x00000010
	NTLMSSP_REQUEST_TARGET                     = 0x00000004
	NTLM_NEGOTIATE_OEM                         = 0x00000002
	NTLMSSP_NEGOTIATE_UNICODE                  = 0x00000001
)

// NegotiateMessage 协商Message
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/b34032e5-3aae-4bc6-84c3-c6d80eadf7f2?source=recommendations
type NegotiateMessage struct {
	Must struct {
		Signature      [8]byte // MUST contain the ASCII string ('N', 'T', 'L', 'M', 'S', 'S', 'P', '\0').
		MessageType    uint32  // This field MUST be set to 0x00000001.
		NegotiateFlags uint32

		DomainName  Field // NTLMSSP_NEGOTIATE_OEM_DOMAIN_SUPPLIED
		Workstation Field // NTLMSSP_NEGOTIATE_OEM_WORKSTATION_SUPPLIED

		Version NVersion // 8bytes
	}
	Optional struct {
		Payload [32]byte
	}
}

func (m *NegotiateMessage) Serialize() []byte {
	if (m.Must.NegotiateFlags & NTLMSSP_NEGOTIATE_VERSION) != 0 {
		m.Must.Version.ProductMajorVersion = WINDOWS_MAJOR_VERSION_6
		m.Must.Version.ProductMinorVersion = WINDOWS_MINOR_VERSION_0
		m.Must.Version.ProductBuild = 6002
		m.Must.Version.NTLMRevisionCurrent = NTLMSSP_REVISION_W2K3
	}
	return core.ToLE(m.Must)
}

func NewNegotiateMessage() *NegotiateMessage {
	m := &NegotiateMessage{}
	m.Must.Signature = [8]byte{'N', 'T', 'L', 'M', 'S', 'S', 'P', 0x00}
	m.Must.MessageType = 0x00000001
	m.Must.NegotiateFlags = NTLMSSP_NEGOTIATE_KEY_EXCH |
		NTLMSSP_NEGOTIATE_128 |
		NTLMSSP_NEGOTIATE_EXTENDED_SESSIONSECURITY |
		NTLMSSP_NEGOTIATE_ALWAYS_SIGN |
		NTLMSSP_NEGOTIATE_NTLM |
		NTLMSSP_NEGOTIATE_SEAL |
		NTLMSSP_NEGOTIATE_SIGN |
		NTLMSSP_REQUEST_TARGET |
		NTLMSSP_NEGOTIATE_UNICODE
	return m
}

func (m *NegotiateMessage) Write(w io.Writer) {
	req := NewTsRequest().SetMessages(m.Serialize())
	req.Write(w)
}

//func (m *NegotiateMessage) Read(r io.Reader) error {
//	treq := &TSRequest{}
//	if data, err := (&core.Asn1{}).Read(r); err != nil {
//		return err
//	} else if _, err := asn1.Unmarshal(data, treq); err != nil {
//		return err
//	} else if len(treq.NegoTokens) == 0 {
//		return fmt.Errorf("invalid TsRequest from upstream, NegoTokens = nil")
//	}
//	return nil
//}
