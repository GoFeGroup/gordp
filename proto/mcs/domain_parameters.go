package mcs

import (
	"bytes"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs/ber"
	"io"
)

type DomainParameters struct {
	MaxChannelIds   int
	MaxUserIds      int
	MaxTokenIds     int
	NumPriorities   int
	MinThoughput    int
	MaxHeight       int
	MaxMCSPDUsize   int
	ProtocolVersion int
}

func (p *DomainParameters) Write(w io.Writer) {
	ber.WriteInteger(w, p.MaxChannelIds)
	ber.WriteInteger(w, p.MaxUserIds)
	ber.WriteInteger(w, p.MaxTokenIds)
	ber.WriteInteger(w, p.NumPriorities)
	ber.WriteInteger(w, p.MinThoughput)
	ber.WriteInteger(w, p.MaxHeight)
	ber.WriteInteger(w, p.MaxMCSPDUsize)
	ber.WriteInteger(w, p.ProtocolVersion)
}

func (p *DomainParameters) Serialize() []byte {
	buff := new(bytes.Buffer)
	p.Write(buff)
	glog.Debugf("DomainParameters: %x", buff.Bytes())
	return buff.Bytes()
}

func (p *DomainParameters) Read(r io.Reader) {
	userData := ber.ReadDomainParameters(r)
	r = bytes.NewReader(userData)
	p.MaxChannelIds = ber.ReadInteger(r)
	p.MaxUserIds = ber.ReadInteger(r)
	p.MaxTokenIds = ber.ReadInteger(r)
	p.NumPriorities = ber.ReadInteger(r)
	p.MinThoughput = ber.ReadInteger(r)
	p.MaxHeight = ber.ReadInteger(r)
	p.MaxMCSPDUsize = ber.ReadInteger(r)
	p.ProtocolVersion = ber.ReadInteger(r)
}
