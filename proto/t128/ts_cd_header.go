package t128

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsCdHeader
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/3d1ccc49-51e2-4a2c-b300-3da9fb931d0e
type TsCdHeader struct {
	CbCompFirstRowSize uint16 // The field MUST be set to 0x0000.
	CbCompMainBodySize uint16 // The size in bytes of the compressed bitmap data
	CbScanWidth        uint16 // The width of the bitmap (which follows this header) in pixels (this value MUST be divisible by 4).
	CbUncompressedSize uint16 // The size in bytes of the bitmap data (which follows this header) after it has been decompressed.
}

func (h *TsCdHeader) Read(r io.Reader) *TsCdHeader {
	core.ReadLE(r, h)
	return h
}
