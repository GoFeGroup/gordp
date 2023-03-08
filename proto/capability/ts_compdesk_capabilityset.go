package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsCompDeskCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/9132002f-f133-4a0f-ba2f-2dc48f1e7f93
type TsCompDeskCapabilitySet struct {
	CompDeskSupportLevel uint16
}

func (c *TsCompDeskCapabilitySet) Type() uint16 {
	return CAPSETTYPE_COMPDESK
}

func (c *TsCompDeskCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsCompDeskCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
