package tpkt

import (
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

// Header --- TPKT Header, as specified in [T123] section 8.
// https://www.itu.int/rec/T-REC-T.123/en
type Header struct {
	Version  uint8
	Reserved uint8
	Length   uint16
}

func (header *Header) Write(w io.Writer) {
	core.WriteBE(w, header)
}

func (header *Header) Read(r io.Reader) {
	core.ReadBE(r, header)
	glog.Debugf("tpkt header: %+v", header)
	core.ThrowIf(header.Version != 3 || header.Length <= 4, fmt.Errorf("invalid tpkt packet"))
}

// Read TPKT Packet data
func Read(r io.Reader) []byte {
	header := &Header{}
	header.Read(r)

	buff := make([]byte, header.Length-4)
	_, err := io.ReadFull(r, buff)
	core.ThrowError(err)
	glog.Debugf("tpkt read header: %v", header)
	glog.Debugf("tpkt read data: %x", buff)
	return buff
}

// Write TPKT packet data
func Write(w io.Writer, data []byte) {
	core.ThrowIf(len(data) > 0xffff-4, fmt.Errorf("invalid data length: %v, can't be more than 0xfffb", len(data)))
	header := &Header{Version: 3, Reserved: 0, Length: uint16(len(data) + 4)}
	header.Write(w)
	glog.Debugf("tpkt write header: %+v", header)
	glog.Debugf("tpkt write data: %x", data)
	core.WriteFull(w, data)
}
