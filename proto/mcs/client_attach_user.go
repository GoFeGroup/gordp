package mcs

import (
	"bytes"
	"io"
)

type ClientAttachUser struct{}

func (c *ClientAttachUser) Write(w io.Writer) {
	WriteMcsPduHeader(w, MCS_PDUTYPE_ATTACH_USER_REQUEST, 0)
}

func (c *ClientAttachUser) Serialize() []byte {
	buff := new(bytes.Buffer)
	c.Write(buff)
	return buff.Bytes()
}
