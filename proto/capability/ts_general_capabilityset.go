package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// OSMajorType
const (
	OSMAJORTYPE_UNSPECIFIED uint16 = 0x0000
	OSMAJORTYPE_WINDOWS            = 0x0001
	OSMAJORTYPE_OS2                = 0x0002
	OSMAJORTYPE_MACINTOSH          = 0x0003
	OSMAJORTYPE_UNIX               = 0x0004
	OSMAJORTYPE_IOS                = 0x0005
	OSMAJORTYPE_OSX                = 0x0006
	OSMAJORTYPE_ANDROID            = 0x0007
)

// OSMinorType
const (
	OSMINORTYPE_UNSPECIFIED    uint16 = 0x0000
	OSMINORTYPE_WINDOWS_31X           = 0x0001
	OSMINORTYPE_WINDOWS_95            = 0x0002
	OSMINORTYPE_WINDOWS_NT            = 0x0003
	OSMINORTYPE_OS2_V21               = 0x0004
	OSMINORTYPE_POWER_PC              = 0x0005
	OSMINORTYPE_MACINTOSH             = 0x0006
	OSMINORTYPE_NATIVE_XSERVER        = 0x0007
	OSMINORTYPE_PSEUDO_XSERVER        = 0x0008
	OSMINORTYPE_WINDOWS_RT            = 0x0009
)
const (
	FASTPATH_OUTPUT_SUPPORTED  uint16 = 0x0001
	NO_BITMAP_COMPRESSION_HDR         = 0x0400
	LONG_CREDENTIALS_SUPPORTED        = 0x0004
	AUTORECONNECT_SUPPORTED           = 0x0008
	ENC_SALTED_CHECKSUM               = 0x0010
)

// TsGeneralCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/41dc6845-07dc-4af6-bc14-d8281acd4877
type TsGeneralCapabilitySet struct {
	OSMajorType             uint16
	OSMinorType             uint16
	ProtocolVersion         uint16 // This field MUST be set to TS_CAPS_PROTOCOLVERSION (0x0200).
	Pad2octetsA             uint16
	GeneralCompressionTypes uint16 // This field MUST be set to zero.
	ExtraFlags              uint16
	UpdateCapabilityFlag    uint16
	RemoteUnshareFlag       uint16
	GeneralCompressionLevel uint16
	RefreshRectSupport      uint8
	SuppressOutputSupport   uint8
}

func (c *TsGeneralCapabilitySet) Type() uint16 {
	return CAPSTYPE_GENERAL
}

func (c *TsGeneralCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsGeneralCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}

func NewTsGeneralCapabilitySet() *TsGeneralCapabilitySet {
	return &TsGeneralCapabilitySet{
		OSMajorType:     OSMAJORTYPE_WINDOWS,
		OSMinorType:     OSMINORTYPE_WINDOWS_NT,
		ProtocolVersion: 0x0200,
		ExtraFlags: LONG_CREDENTIALS_SUPPORTED | NO_BITMAP_COMPRESSION_HDR |
			ENC_SALTED_CHECKSUM | FASTPATH_OUTPUT_SUPPORTED,
	}
}
