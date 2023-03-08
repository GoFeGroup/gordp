package t128

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

const (
	PDUTYPE2_UPDATE                      = 0x02
	PDUTYPE2_CONTROL                     = 0x14
	PDUTYPE2_POINTER                     = 0x1B
	PDUTYPE2_INPUT                       = 0x1C
	PDUTYPE2_SYNCHRONIZE                 = 0x1F
	PDUTYPE2_REFRESH_RECT                = 0x21
	PDUTYPE2_PLAY_SOUND                  = 0x22
	PDUTYPE2_SUPPRESS_OUTPUT             = 0x23
	PDUTYPE2_SHUTDOWN_REQUEST            = 0x24
	PDUTYPE2_SHUTDOWN_DENIED             = 0x25
	PDUTYPE2_SAVE_SESSION_INFO           = 0x26
	PDUTYPE2_FONTLIST                    = 0x27
	PDUTYPE2_FONTMAP                     = 0x28
	PDUTYPE2_SET_KEYBOARD_INDICATORS     = 0x29
	PDUTYPE2_BITMAPCACHE_PERSISTENT_LIST = 0x2B
	PDUTYPE2_BITMAPCACHE_ERROR_PDU       = 0x2C
	PDUTYPE2_SET_KEYBOARD_IME_STATUS     = 0x2D
	PDUTYPE2_OFFSCRCACHE_ERROR_PDU       = 0x2E
	PDUTYPE2_SET_ERROR_INFO_PDU          = 0x2F
	PDUTYPE2_DRAWNINEGRID_ERROR_PDU      = 0x30
	PDUTYPE2_DRAWGDIPLUS_ERROR_PDU       = 0x31
	PDUTYPE2_ARC_STATUS_PDU              = 0x32
	PDUTYPE2_STATUS_INFO_PDU             = 0x36
	PDUTYPE2_MONITOR_LAYOUT_PDU          = 0x37
)

// StreamId
const (
	STREAM_UNDEFINED = 0x00
	STREAM_LOW       = 0x01
	STREAM_MED       = 0x02
	STREAM_HI        = 0x04
)

// Level-2 Compression Flags
const (
	PACKET_COMPRESSED = 0x20
	PACKET_AT_FRONT   = 0x40
	PACKET_FLUSHED    = 0x80
)

// Level-1 Compression Flags
const (
	L1_PACKET_AT_FRONT   = 0x04
	L1_NO_COMPRESSION    = 0x02
	L1_COMPRESSED        = 0x01
	L1_INNER_COMPRESSION = 0x10
)

const (
	PACKET_COMPR_TYPE_8K    = 0x0
	PACKET_COMPR_TYPE_64K   = 0x1
	PACKET_COMPR_TYPE_RDP6  = 0x2
	PACKET_COMPR_TYPE_RDP61 = 0x3
)

// TsShareDataHeader
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/4b5d4c0d-a657-41e9-9c69-d58632f46d31
type TsShareDataHeader struct {
	SharedId           uint32
	Padding1           uint8
	StreamId           uint8
	UncompressedLength uint16
	PDUType2           uint8

	//https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/9355a663-ef22-4431-afeb-d72ac68f25fd
	CompressedType   uint8
	CompressedLength uint16
}

func (h *TsShareDataHeader) Read(r io.Reader) {
	core.ReadLE(r, h)
	glog.Debugf("[!] compressedType: %x", h.CompressedType)
	core.ThrowIf(h.CompressedType&PACKET_COMPRESSED != 0, "compress not implement")
}
