package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsControlCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/e0add8ac-1546-4091-85ba-0ea77f54f2c7
type TsControlCapabilitySet struct {
	ControlFlags     uint16
	RemoteDetachFlag uint16
	ControlInterest  uint16
	DetachInterest   uint16
}

func (c *TsControlCapabilitySet) Type() uint16 {
	return CAPSTYPE_CONTROL
}

func (c *TsControlCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsControlCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
