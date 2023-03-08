// Package nla --- CredCSSP 协议
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cssp/85f57821-40bb-46aa-bfcb-ba9590b8fc30
package nla

import (
	"encoding/asn1"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

// NegoToken
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cssp/9664994d-0784-4659-b85b-83b8d54c2336
type NegoToken struct {
	Data []byte `asn1:"explicit,tag:0"`
}

// TSCredentials
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cssp/94a1ab00-5500-42fd-8d3d-7a84e6c2cf03
type TSCredentials struct {
	CredType    int    `asn1:"explicit,tag:0"`
	Credentials []byte `asn1:"explicit,tag:1"`
}

func (c TSCredentials) Serialize() []byte {
	data, _ := asn1.Marshal(c)
	return data
}

// TSPasswordCreds
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cssp/17773cc4-21e9-4a75-a0dd-72706b174fe5
type TSPasswordCreds struct {
	DomainName []byte `asn1:"explicit,tag:0"`
	UserName   []byte `asn1:"explicit,tag:1"`
	Password   []byte `asn1:"explicit,tag:2"`
}

func (c TSPasswordCreds) Serialize() []byte {
	data, _ := asn1.Marshal(c)
	return data
}

// TSCspDataDetail
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cssp/34ee27b3-5791-43bb-9201-076054b58123
type TSCspDataDetail struct {
	KeySpec       int    `asn1:"explicit,tag:0"`
	CardName      string `asn1:"explicit,tag:1"`
	ReaderName    string `asn1:"explicit,tag:2"`
	ContainerName string `asn1:"explicit,tag:3"`
	CspName       string `asn1:"explicit,tag:4"`
}

// TSSmartCardCreds
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cssp/4251d165-cf01-4513-a5d8-39ee4a98b7a4
type TSSmartCardCreds struct {
	Pin        string            `asn1:"explicit,tag:0"`
	CspData    []TSCspDataDetail `asn1:"explicit,tag:1"`
	UserHint   string            `asn1:"explicit,tag:2"`
	DomainHint string            `asn1:"explicit,tag:3"`
}

// TSRequest
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cssp/6aac4dea-08ef-47a6-8747-22ea7f6d8685
type TSRequest struct {
	Version     int         `asn1:"explicit,tag:0"`
	NegoTokens  []NegoToken `asn1:"optional,explicit,tag:1"`
	AuthInfo    []byte      `asn1:"optional,explicit,tag:2"`
	PubKeyAuth  []byte      `asn1:"optional,explicit,tag:3"`
	ErrorCode   int         `asn1:"optional,explicit,tag:4"`
	ClientNonce int         `asn1:"optional,explicit,tag:5"`
}

func NewTsRequest() *TSRequest {
	return &TSRequest{Version: 2}
}

func (req *TSRequest) SetMessages(data []byte) *TSRequest {
	token := NegoToken{data}
	req.NegoTokens = append(req.NegoTokens, token)
	return req
}

func (req *TSRequest) SetAuthInfo(authInfo []byte) *TSRequest {
	req.AuthInfo = core.If(len(authInfo) > 0, authInfo, nil)
	return req
}

func (req *TSRequest) SetPubKeyAuth(pubKeyAuth []byte) *TSRequest {
	req.PubKeyAuth = core.If(len(pubKeyAuth) > 0, pubKeyAuth, nil)
	return req
}

func (req *TSRequest) Read(r io.Reader) {
	data := (&core.Asn1{}).Read(r)
	_, err := asn1.Unmarshal(data, req)
	core.ThrowError(err)
	glog.Debugf("TsRequest Read: %x", data)
}

func (req *TSRequest) Write(w io.Writer) {
	data, _ := asn1.Marshal(*req)
	glog.Debugf("TsRequest Write: %v, %x", len(data), data)
	core.WriteFull(w, data)
}

//func EncodeDERTRequest(msgs []Message, authInfo []byte, pubKeyAuth []byte) ([]byte, error) {
//	req := (&TSRequest{Version: 2}).SetMessages(msgs).SetAuthInfo(authInfo).SetPubKeyAuth(pubKeyAuth)
//	return asn1.Marshal(*req)
//}
//
//func DecodeDERTRequest(s []byte) (*TSRequest, error) {
//	treq := &TSRequest{}
//	_, err := asn1.Unmarshal(s, treq)
//	return treq, err
//}
//
//func EncodeDERTCredentials(domain, username, password []byte) ([]byte, error) {
//	if tpass, err := asn1.Marshal(TSPasswordCreds{domain, username, password}); err != nil {
//		return nil, err
//	} else {
//		return asn1.Marshal(TSCredentials{1, tpass})
//	}
//}
//
//func DecodeDERTCredentials(s []byte) (*TSCredentials, error) {
//	tcre := &TSCredentials{}
//	_, err := asn1.Unmarshal(s, tcre)
//	return tcre, err
//}
