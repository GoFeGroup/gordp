package per

import (
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

// Package per
// FIXME: Copy From FreeRDP C.gcc_write_conference_create_request
func fixme() {}

func WriteChoice(w io.Writer, choice byte) {
	core.WriteFull(w, []byte{choice})
}

func ReadChoice(r io.Reader) byte {
	return ReadInteger8(r)
}

func WriteObjectIdentifier(w io.Writer, oid []byte) {
	t12 := oid[0]*40 + oid[1]
	core.WriteFull(w, []byte{0x5}) // length
	core.WriteFull(w, []byte{t12}) // first two tuples
	core.WriteFull(w, oid[2:6])
}

func ReadObjectIdentifier(r io.Reader) []byte {
	length := ReadLength(r)
	core.ThrowIf(length != 5, fmt.Errorf("invalid length of ObjectIdentifier: %v != 5", length))
	oid := make([]byte, 6)
	t12 := ReadInteger8(r) // first two tuple
	oid[0] = t12 / 40
	oid[1] = t12 % 40
	core.ReadBE(r, oid[2:6])
	return oid
}

func WriteLength(w io.Writer, length int) {
	if length > 0x7F {
		core.WriteBE(w, uint16(length|0x8000))
	} else {
		core.WriteFull(w, []byte{uint8(length & 0xff)})
	}
}

func ReadLength(r io.Reader) int {
	b := ReadInteger8(r)
	if b&0x80 != 0 {
		length := int(b&^0x80) << 8
		return length + int(ReadInteger8(r))
	}
	return int(b)
}

func WriteSelection(w io.Writer, selection uint8) {
	core.WriteFull(w, []byte{selection})
}

func WriteNumericString(w io.Writer, numStr string, minValue int) {
	length := len(numStr)
	mLen := core.If(length >= minValue, length-minValue, minValue)

	WriteLength(w, mLen)
	for i := 0; i < length; i += 2 {
		c1 := numStr[i]
		c2 := 0x30
		if i+1 < length {
			c2 = int(numStr[i+1])
		}
		//c2 := core.If((i+1) < length, numStr[i+1], 0x30)
		c1 = (c1 - 0x30) % 10
		c2 = (c2 - 0x30) % 10
		num := (c1 << 4) | uint8(c2)
		core.WriteFull(w, []byte{num})
	}
}

func WritePadding(w io.Writer, padLen int) {
	core.WriteFull(w, make([]byte, padLen))
}

func WriteNumberOfSet(w io.Writer, n int) {
	core.WriteFull(w, []byte{uint8(n)})
}

func ReadNumberOfSet(r io.Reader) int {
	return int(ReadInteger8(r))
}

func WriteOctetString(w io.Writer, oStr string, minValue int) {
	length := len(oStr)
	mLen := core.If(length >= minValue, length-minValue, minValue)

	WriteLength(w, mLen)
	core.WriteFull(w, []byte(oStr))
}

func ReadOctetString(r io.Reader, minValue int) []byte {
	length := ReadLength(r)
	buff := make([]byte, length+minValue)
	_, err := io.ReadFull(r, buff)
	core.ThrowError(err)
	return buff
}

func ReadInteger8(r io.Reader) (length uint8) {
	core.ReadBE(r, &length)
	return length
}

func WriteInteger8(w io.Writer, n uint8) {
	core.WriteBE(w, n)
}

func ReadInteger16(r io.Reader, min uint16) uint16 {
	var i16 uint16
	core.ReadBE(r, &i16)
	if i16 > 0xffff-min {
		glog.Warnf("PER uint16 invalid value %0#x > %0#x", i16, 0xffff-min)
	}
	return i16 + min
}

func WriteInteger16(w io.Writer, n uint16) {
	core.WriteBE(w, n)
}

func WriteInteger32(w io.Writer, n uint32) {
	core.WriteBE(w, n)
}

func ReadInteger(r io.Reader) uint32 {
	length := ReadLength(r)
	if length == 0 {
		return 0
	} else if length == 1 {
		return uint32(ReadInteger8(r))
	} else if length == 2 {
		return uint32(ReadInteger16(r, 0))
	}
	core.Throw(fmt.Errorf("invalid length of integer: %v", length))
	return 0
}

func WriteInteger(w io.Writer, n uint32) {
	if n <= 0xFF {
		WriteLength(w, 1)
		WriteInteger8(w, uint8(n))
	} else if n <= 0xFFFF {
		WriteLength(w, 2)
		WriteInteger16(w, uint16(n))
	} else {
		WriteLength(w, 4)
		WriteInteger32(w, n)
	}
}

func ReadEnumerated(r io.Reader) uint8 {
	return ReadInteger8(r)
}
