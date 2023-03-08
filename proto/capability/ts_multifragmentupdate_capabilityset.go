package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsMultiFragmentUpdateCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/01717954-716a-424d-af35-28fb2b86df89
type TsMultiFragmentUpdateCapabilitySet struct {
	MaxRequestSize uint32
}

func (c *TsMultiFragmentUpdateCapabilitySet) Type() uint16 {
	return CAPSETTYPE_MULTIFRAGMENTUPDATE
}

func (c *TsMultiFragmentUpdateCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsMultiFragmentUpdateCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
