package nla

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"io"
)

const (
	MsvAvEOL             = 0x0000
	MsvAvNbComputerName  = 0x0001
	MsvAvNbDomainName    = 0x0002
	MsvAvDnsComputerName = 0x0003
	MsvAvDnsDomainName   = 0x0004
	MsvAvDnsTreeName     = 0x0005
	MsvAvFlags           = 0x0006
	MsvAvTimestamp       = 0x0007
	MsvAvSingleHost      = 0x0008
	MsvAvTargetName      = 0x0009
	MsvChannelBindings   = 0x000A
)

// AVPair
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/83f5e789-660d-4781-8491-5f8c6641f75e
type AVPair struct {
	Must struct {
		Id  uint16
		Len uint16
	}

	Optional struct {
		Value []byte
	}
}

func (avPair *AVPair) Read(r io.Reader) {
	core.ReadLE(r, &avPair.Must)
	avPair.Optional.Value = make([]byte, avPair.Must.Len)
	_, err := io.ReadFull(r, avPair.Optional.Value)
	core.ThrowError(err)
}

func (avPair *AVPair) Write(w io.Writer) {
	core.WriteLE(w, &avPair.Must)
	if avPair.Must.Len > 0 {
		core.WriteFull(w, avPair.Optional.Value)
	}
}

type AVPairs []AVPair

func (avPairs AVPairs) Write(w io.Writer) {
	for _, v := range avPairs {
		v.Write(w)
	}
}

func ReadAvPairs(data []byte) AVPairs {
	var avPairs AVPairs
	r := bytes.NewReader(data)
	for r.Len() > 0 {
		avPair := AVPair{}
		avPair.Read(r)
		//glog.Debugf("avPair: id: %v, len: %v", avPair.Must.Id, avPair.Must.Len)
		avPairs = append(avPairs, avPair)
		if avPair.Must.Id == MsvAvEOL {
			break
		}
	}
	return avPairs
}

func (avPairs AVPairs) GetTimeStamp() []byte {
	for _, v := range avPairs {
		switch v.Must.Id {
		case MsvAvTimestamp:
			return v.Optional.Value
		case MsvAvEOL:
			return nil
		}
	}
	return nil
}
