package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsSurfCmdsCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/aa953018-c0a8-4761-bb12-86586c2cd56a
type TsSurfCmdsCapabilitySet struct {
	CmdFlags uint32
	Reserved uint32
}

func (c *TsSurfCmdsCapabilitySet) Type() uint16 {
	return CAPSETTYPE_SURFACE_COMMANDS
}

func (c *TsSurfCmdsCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsSurfCmdsCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
