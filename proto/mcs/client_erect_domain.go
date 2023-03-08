package mcs

import (
	"bytes"
	"github.com/GoFeGroup/gordp/proto/mcs/per"
	"io"
)

type ClientErectDomain struct{}

func (e *ClientErectDomain) Write(w io.Writer) {
	WriteMcsPduHeader(w, MCS_PDUTYPE_ERECT_DOMAIN_REQUEST, 0)
	per.WriteInteger(w, 0) // subHeight
	per.WriteInteger(w, 0)
}

func (e *ClientErectDomain) Serialize() []byte {
	buff := new(bytes.Buffer)
	e.Write(buff)
	return buff.Bytes()
}
