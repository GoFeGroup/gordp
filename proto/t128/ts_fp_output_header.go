package t128

import (
	"io"

	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
)

// update code
const (
	FASTPATH_UPDATETYPE_ORDERS        = 0x0
	FASTPATH_UPDATETYPE_BITMAP        = 0x1
	FASTPATH_UPDATETYPE_PALETTE       = 0x2
	FASTPATH_UPDATETYPE_SYNCHRONIZE   = 0x3
	FASTPATH_UPDATETYPE_SURFCMDS      = 0x4
	FASTPATH_UPDATETYPE_PTR_NULL      = 0x5
	FASTPATH_UPDATETYPE_PTR_DEFAULT   = 0x6
	FASTPATH_UPDATETYPE_PTR_POSITION  = 0x8
	FASTPATH_UPDATETYPE_COLOR         = 0x9
	FASTPATH_UPDATETYPE_CACHED        = 0xA
	FASTPATH_UPDATETYPE_POINTER       = 0xB
	FASTPATH_UPDATETYPE_LARGE_POINTER = 0xC
)

// fragmentation
const (
	FASTPATH_FRAGMENT_SINGLE = 0x0
	FASTPATH_FRAGMENT_LAST   = 0x1
	FASTPATH_FRAGMENT_FIRST  = 0x2
	FASTPATH_FRAGMENT_NEXT   = 0x3
)

// compression
const (
	FASTPATH_OUTPUT_COMPRESSION_USED = 0x2
)

// FpOutputHeader
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/a1c4caa8-00ed-45bb-a06e-5177473766d3
type FpOutputHeader struct {
	UpdateCode    uint8
	Fragmentation uint8
	Compression   uint8
}

func (h *FpOutputHeader) Read(r io.Reader) {
	var updateHeader uint8
	core.ReadLE(r, &updateHeader)
	glog.Debugf("fpOutputHeader: %x", updateHeader)
	h.UpdateCode = updateHeader & 0xF
	h.Fragmentation = (updateHeader >> 4) & 0x03
	h.Compression = (updateHeader >> 6) & 0x03
	glog.Debugf("fpOutputHeader: %+v", h)

	if h.Compression == FASTPATH_OUTPUT_COMPRESSION_USED {
		core.Throw("not implement")
	}
}
