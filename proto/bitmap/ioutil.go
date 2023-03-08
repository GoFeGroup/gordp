package bitmap

import (
	"encoding/binary"
	"github.com/GoFeGroup/gordp/core"
	"io"
)

func ReadByte(r io.Reader) byte {
	var b byte
	core.ThrowError(binary.Read(r, binary.LittleEndian, &b))
	return b
}

func WriteByte(w io.Writer, b byte) {
	core.ThrowError(binary.Write(w, binary.LittleEndian, b))
}

func ReadBytes(r io.Reader, length int) []byte {
	b := make([]byte, length)
	_, err := io.ReadFull(r, b)
	core.ThrowError(err)
	return b
}

func WriteBytes(w io.Writer, data []byte) {
	core.ThrowError(binary.Write(w, binary.LittleEndian, data))
}

func ReadShortLE(r io.Reader) uint16 {
	var b uint16
	core.ThrowError(binary.Read(r, binary.LittleEndian, &b))
	return b
}

func WriteShortLE(w io.Writer, b uint16) {
	core.ThrowError(binary.Write(w, binary.LittleEndian, b))
}

func ReadIntLE(r io.Reader) uint32 {
	var b uint32
	core.ThrowError(binary.Read(r, binary.LittleEndian, &b))
	return b
}
