package nla

// NTLMv2Response
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/d43e2224-6fc3-449d-9f37-b90b55a29c80
type NTLMv2Response struct {
	Response              [16]byte
	NTLMv2ClientChallenge []byte
}
