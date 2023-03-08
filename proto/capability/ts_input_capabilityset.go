package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"io"
)

// http://msdn.microsoft.com/en-us/library/cc240563.aspx
const (
	INPUT_FLAG_SCANCODES       uint16 = 0x0001
	INPUT_FLAG_MOUSEX                 = 0x0004
	INPUT_FLAG_FASTPATH_INPUT         = 0x0008
	INPUT_FLAG_UNICODE                = 0x0010
	INPUT_FLAG_FASTPATH_INPUT2        = 0x0020
	INPUT_FLAG_UNUSED1                = 0x0040
	INPUT_FLAG_UNUSED2                = 0x0080
	INPUT_FLAG_MOUSE_HWHEEL           = 0x0100
)

// TsInputCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/b3bc76ae-9ee5-454f-b197-ede845ca69cc
type TsInputCapabilitySet struct {
	Flags               uint16
	Pad2octetsA         uint16
	KeyboardLayout      uint32
	KeyboardType        uint32
	KeyboardSubType     uint32
	KeyboardFunctionKey uint32
	ImeFileName         [64]byte
}

func (c *TsInputCapabilitySet) Type() uint16 {
	return CAPSTYPE_INPUT
}

func (c *TsInputCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsInputCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
	core.WriteLE(w, []byte{0x0c, 0, 0, 0})
}

func NewTsInputCapabilitySet() *TsInputCapabilitySet {
	return &TsInputCapabilitySet{
		Flags:               INPUT_FLAG_SCANCODES | INPUT_FLAG_MOUSEX | INPUT_FLAG_UNICODE,
		KeyboardLayout:      mcs.US,
		KeyboardType:        mcs.KT_IBM_101_102_KEYS,
		KeyboardSubType:     0,
		KeyboardFunctionKey: 12,
		ImeFileName:         [64]byte{},
	}
}
