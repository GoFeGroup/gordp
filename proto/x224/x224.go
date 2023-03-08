package x224

import (
	"bytes"
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/tpkt"
	"io"
)

const (
	TPDU_CONNECTION_REQUEST = 0xE0
	TPDU_CONNECTION_CONFIRM = 0xD0
	TPDU_DISCONNECT_REQUEST = 0x80
	TPDU_DATA               = 0xF0
	TPDU_ERROR              = 0x70
)

// Header --- An X.224 Class 0 Connection Request transport protocol data unit (TPDU), as specified in [X224] section 13.3.
// https://www.itu.int/rec/T-REC-X.224-199511-I/en
type Header struct {
	Length  uint8
	PduType uint8
	DstRef  uint16
	SrcRef  uint16
	Flags   uint8
}

func (header *Header) Write(w io.Writer) {
	core.WriteLE(w, header)
}

func (header *Header) Read(r io.Reader) {
	core.ReadLE(r, header)
}

func Connect(w io.Writer, pduType uint8, data []byte) {
	core.ThrowIf(len(data) > 0xf9, fmt.Errorf("invalid data length: %v, can't be more than 0xf9", len(data)))
	header := &Header{uint8(6 + len(data)), pduType, 0, 0, 0}
	buf := new(bytes.Buffer)
	header.Write(buf)
	core.WriteFull(buf, data)
	glog.Debugf("x224 write: %x", buf.Bytes())
	tpkt.Write(w, buf.Bytes())
}

// ReadConfirm TPDU type and data
func ReadConfirm(r io.Reader) (uint8, []byte) {
	data := tpkt.Read(r)
	glog.Debugf("x224 read: %x", data)
	buf := bytes.NewBuffer(data)
	header := &Header{}
	header.Read(buf)
	core.ThrowIf(int(header.Length) != len(data)-1, fmt.Errorf("invalid x224 length"))
	return header.PduType, data[7:]
}

func Write(w io.Writer, data []byte) {
	var arr = [][]byte{{2, TPDU_DATA, 0x80}, data}
	glog.Debugf("x224 write %v - %x", len(data)+3, bytes.Join(arr, nil))
	tpkt.Write(w, bytes.Join(arr, nil))
}

func Read(r io.Reader) []byte {
	data := tpkt.Read(r)
	if len(data) <= 3 || !bytes.Equal(data[:3], []byte{2, TPDU_DATA, 0x80}) {
		core.Throw(fmt.Errorf("invalid x224 header"))
	}
	glog.Debugf("x224 read %v - %x", len(data[3:]), data[3:])
	return data[3:]
}
