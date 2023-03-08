package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsWindowActivationCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/97ff3178-9999-4231-ae4c-1e8d10d0e219
type TsWindowActivationCapabilitySet struct {
	HelpKeyFlag          uint16
	HelpKeyIndexFlag     uint16
	HelpExtendedKeyFlag  uint16
	WindowManagerKeyFlag uint16
}

func (c *TsWindowActivationCapabilitySet) Type() uint16 {
	return CAPSTYPE_ACTIVATION
}

func (c *TsWindowActivationCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsWindowActivationCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
