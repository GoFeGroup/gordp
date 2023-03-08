package ber

import (
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

// FIXME: Copy from Freerdp C.mcs_write_connect_initial
func fixme() {}

// Class - bits 8 and 7
const (
	BER_CLASS_MASK = 0xC0
	BER_CLASS_UNIV = 0x00 /* 0 0 */
	BER_CLASS_APPL = 0x40 /* 0 1 */
	BER_CLASS_CTXT = 0x80 /* 1 0 */
	BER_CLASS_PRIV = 0xC0 /* 1 1 */
)

// P/C - bit 6
const (
	BER_PC_MASK   = 0x20
	BER_PRIMITIVE = 0x00 /* 0 */
	BER_CONSTRUCT = 0x20 /* 1 */
)

const (
	BER_TAG_MASK            = 0x1F
	BER_TAG_BOOLEAN         = 0x01
	BER_TAG_INTEGER         = 0x02
	BER_TAG_BIT_STRING      = 0x03
	BER_TAG_OCTET_STRING    = 0x04
	BER_TAG_OBJECT_IDENFIER = 0x06
	BER_TAG_ENUMERATED      = 0x0A
	BER_TAG_SEQUENCE        = 0x10
	BER_TAG_SEQUENCE_OF     = 0x10
)

func BER_PC(pc bool) uint8 {
	return uint8(core.If(pc, BER_CONSTRUCT, BER_PRIMITIVE))
}

func WriteUniversalTag(w io.Writer, tag byte, pc bool) {
	core.WriteBE(w, []byte{(BER_CLASS_UNIV | BER_PC(pc)) | (BER_TAG_MASK & tag)})
}

func ReadUniversalTag(r io.Reader, tag byte, pc bool) bool {
	bb := ReadInteger8(r)
	return bb == (BER_CLASS_UNIV|BER_PC(pc))|(BER_TAG_MASK&tag)
}

func WriteLength(w io.Writer, length int) {
	if length > 0xff {
		core.WriteBE(w, []byte{0x80 ^ 2})
		core.WriteBE(w, uint16(length))
	} else if length > 0x7f {
		core.WriteBE(w, []byte{0x80 ^ 1})
		core.WriteBE(w, uint8(length))
	} else {
		core.WriteBE(w, uint8(length))
	}
}

func ReadLength(r io.Reader) int {
	size := ReadInteger8(r)
	if size&0x80 == 0 {
		return int(size)
	}
	switch size = size &^ 0x80; size {
	case 1:
		return int(ReadInteger8(r))
	case 2:
		return int(ReadInteger16(r))
	}
	core.Throw(fmt.Errorf("can not reach here"))
	return 0
}

func WriteInteger(w io.Writer, n int) {
	WriteUniversalTag(w, BER_TAG_INTEGER, false)
	if n <= 0xff {
		WriteLength(w, 1)
		core.WriteBE(w, uint8(n&0xff))
	} else if n <= 0xffff {
		WriteLength(w, 2)
		core.WriteBE(w, uint16(n&0xffff))
	} else {
		WriteLength(w, 4)
		core.WriteBE(w, uint32(n))
	}
}

func WriteBoolean(w io.Writer, ok bool) {
	WriteUniversalTag(w, BER_TAG_BOOLEAN, false)
	WriteLength(w, 1)
	core.WriteBE(w, uint8(core.If(ok, 0xff, 0)))
}

func WriteOctetstring(w io.Writer, str string) {
	WriteUniversalTag(w, BER_TAG_OCTET_STRING, false)
	WriteLength(w, len(str))
	core.WriteFull(w, []byte(str))
}

func ReadDomainParameters(r io.Reader) []byte {
	core.ThrowIf(ReadUniversalTag(r, BER_TAG_SEQUENCE, true) == false, "invalid universal tag")
	length := ReadLength(r)
	buff := make([]byte, length)
	_, err := io.ReadFull(r, buff)
	core.ThrowError(err)
	return buff
}

func WriteDomainParameters(w io.Writer, data []byte) {
	WriteUniversalTag(w, BER_TAG_SEQUENCE, true)
	WriteLength(w, len(data))
	core.WriteFull(w, data)
}

func WriteApplicationTag(w io.Writer, tag uint8, data []byte) {
	if tag > 30 {
		core.WriteBE(w, []byte{BER_CLASS_APPL | BER_CONSTRUCT | BER_TAG_MASK, tag})
		WriteLength(w, len(data))
	} else {
		core.WriteBE(w, []byte{(BER_CLASS_APPL | BER_CONSTRUCT) | (BER_TAG_MASK & tag)})
		WriteLength(w, len(data))
	}
	core.WriteFull(w, data)
}

func ReadApplicationTag(r io.Reader, tag uint8) []byte {
	bb := ReadInteger8(r)
	if tag > 30 {
		core.ThrowIf(bb != (BER_CLASS_APPL|BER_CONSTRUCT)|BER_TAG_MASK, "ReadApplicationTag: invalud data")
		if bb := ReadInteger8(r); bb != tag {
			core.Throw(fmt.Errorf("ReadApplicationTag: invalid tag, need %v, but it's %v", tag, bb))
		}
	} else if bb != (BER_CLASS_MASK|BER_CONSTRUCT)|(BER_TAG_MASK&tag) {
		core.Throw(fmt.Errorf("ReadApplicationTag: invalid tag, need %v, but it's %v", tag, bb))
	}

	length := ReadLength(r)
	glog.Debugf("read application tag: %v", length)
	return core.ReadBytes(r, length)
}

func ReadInteger8(r io.Reader) (bb uint8) {
	core.ReadBE(r, &bb)
	return
}

func ReadInteger16(r io.Reader) (bb uint16) {
	core.ReadBE(r, &bb)
	return
}

func ReadInteger32(r io.Reader) (bb uint32) {
	core.ReadBE(r, &bb)
	return
}

func ReadInteger(r io.Reader) int {
	core.ThrowIf(ReadUniversalTag(r, BER_TAG_INTEGER, false) == false, "invalid integer tag")
	switch length := ReadLength(r); length {
	case 1:
		return int(ReadInteger8(r))
	case 2:
		return int(ReadInteger16(r))
	case 3:
		l1 := ReadInteger8(r)
		l2 := ReadInteger16(r)
		return int(l1)<<16 + int(l2)
	case 4:
		return int(ReadInteger32(r))
	}
	core.Throw("can't reach here")
	return 0
}

func ReadEnumerated(r io.Reader) uint8 {
	core.ThrowIf(ReadUniversalTag(r, BER_TAG_ENUMERATED, false) == false, "invalid enumerated tag")
	length := ReadLength(r)
	core.ThrowIf(length != 1, fmt.Errorf("invalid length %v, not 1", length))
	return ReadInteger8(r)
}
