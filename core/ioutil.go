package core

import (
	"bytes"
	"encoding/binary"
	"io"
)

func ReadBytes(r io.Reader, length int) []byte {
	data := make([]byte, length)
	return ReadFull(r, data)
}

func ReadFull(r io.Reader, data []byte) []byte {
	for pos := 0; pos < len(data); {
		n, err := io.ReadFull(r, data[pos:])
		ThrowError(err)
		pos += n
	}
	return data
}

func WriteFull(w io.Writer, data []byte) {
	for pos := 0; pos < len(data); {
		n, err := w.Write(data[pos:])
		ThrowIf(err != nil && err != io.ErrShortWrite, err)
		pos += n
	}
}

func ReadBE[T any](r io.Reader, data T) T {
	ThrowError(binary.Read(r, binary.BigEndian, data))
	return data
}

func ReadLE[T any](r io.Reader, data T) T {
	ThrowError(binary.Read(r, binary.LittleEndian, data))
	return data
}

func WriteBE[T any](w io.Writer, data T) {
	ThrowError(binary.Write(w, binary.BigEndian, data))
}

func WriteLE[T any](w io.Writer, data T) {
	ThrowError(binary.Write(w, binary.LittleEndian, data))
}

func ToBE[T any](o T) []byte {
	buff := new(bytes.Buffer)
	WriteBE(buff, o)
	return buff.Bytes()
}

func ToLE[T any](o T) []byte {
	buff := new(bytes.Buffer)
	WriteLE(buff, o)
	return buff.Bytes()
}
