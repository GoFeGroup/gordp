package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsVirtualChannelCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/a8593178-80c0-4b80-876c-cb77e62cecfc
type TsVirtualChannelCapabilitySet struct {
	Flags       uint32
	VCChunkSize uint32
}

func (c *TsVirtualChannelCapabilitySet) Type() uint16 {
	return CAPSTYPE_VIRTUALCHANNEL
}

func (c *TsVirtualChannelCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsVirtualChannelCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
