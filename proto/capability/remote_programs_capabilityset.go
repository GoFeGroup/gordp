package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

const (
	RAIL_LEVEL_SUPPORTED                           = 0x00000001
	RAIL_LEVEL_DOCKED_LANGBAR_SUPPORTED            = 0x00000002
	RAIL_LEVEL_SHELL_INTEGRATION_SUPPORTED         = 0x00000004
	RAIL_LEVEL_LANGUAGE_IME_SYNC_SUPPORTED         = 0x00000008
	RAIL_LEVEL_SERVER_TO_CLIENT_IME_SYNC_SUPPORTED = 0x00000010
	RAIL_LEVEL_HIDE_MINIMIZED_APPS_SUPPORTED       = 0x00000020
	RAIL_LEVEL_WINDOW_CLOAKING_SUPPORTED           = 0x00000040
	RAIL_LEVEL_HANDSHAKE_EX_SUPPORTED              = 0x00000080
)

// RemoteProgramsCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdperp/36a25e21-25e1-4954-aae8-09aaf6715c79
type RemoteProgramsCapabilitySet struct {
	RailSupportLevel uint32
}

func (c *RemoteProgramsCapabilitySet) Type() uint16 {
	return CAPSTYPE_RAIL
}

func (c *RemoteProgramsCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *RemoteProgramsCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}

func NewRemoteProgramsCapabilitySet() *RemoteProgramsCapabilitySet {
	return &RemoteProgramsCapabilitySet{
		RailSupportLevel: RAIL_LEVEL_SUPPORTED |
			RAIL_LEVEL_SHELL_INTEGRATION_SUPPORTED |
			RAIL_LEVEL_LANGUAGE_IME_SYNC_SUPPORTED |
			RAIL_LEVEL_SERVER_TO_CLIENT_IME_SYNC_SUPPORTED |
			RAIL_LEVEL_HIDE_MINIMIZED_APPS_SUPPORTED |
			RAIL_LEVEL_WINDOW_CLOAKING_SUPPORTED |
			RAIL_LEVEL_HANDSHAKE_EX_SUPPORTED |
			RAIL_LEVEL_DOCKED_LANGBAR_SUPPORTED,
	}
}
