package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

const (
	NEGOTIATEORDERSUPPORT   = 0x0002
	ZEROBOUNDSDELTASSUPPORT = 0x0008
	COLORINDEXSUPPORT       = 0x0020
	SOLIDPATTERNBRUSHONLY   = 0x0040
	ORDERFLAGS_EXTRA_FLAGS  = 0x0080
)

// TsOrderCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/9f409c29-480c-4751-9665-510b8ffff294
type TsOrderCapabilitySet struct {
	TerminalDescriptor      [16]byte
	Pad4octetsA             uint32
	DesktopSaveXGranularity uint16
	DesktopSaveYGranularity uint16
	Pad2octetsA             uint16
	MaximumOrderLevel       uint16
	NumberFonts             uint16
	OrderFlags              uint16
	OrderSupport            [32]byte
	TextFlags               uint16
	OrderSupportExFlags     uint16
	Pad4octetsB             uint32
	DesktopSaveSize         uint32
	Pad2octetsC             uint16
	Pad2octetsD             uint16
	TextANSICodePage        uint16
	Pad2octetsE             uint16
}

func (c *TsOrderCapabilitySet) Type() uint16 {
	return CAPSTYPE_ORDER
}

func (c *TsOrderCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsOrderCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}

func NewTsOrderCapabilitySet() *TsOrderCapabilitySet {
	return &TsOrderCapabilitySet{
		DesktopSaveXGranularity: 1,
		DesktopSaveYGranularity: 20,
		MaximumOrderLevel:       1,
		OrderFlags:              NEGOTIATEORDERSUPPORT | ZEROBOUNDSDELTASSUPPORT,
		DesktopSaveSize:         480 * 480,
	}
}
