package secPdu

// SecurityExchangePDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/ca73831d-3661-4700-9357-8f247640c02e
type SecurityExchangePDU struct {
	BasicSecurityHeader   uint32
	Length                uint32
	EncryptedClientRandom []byte
}
