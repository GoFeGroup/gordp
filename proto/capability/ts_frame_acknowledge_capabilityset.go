package capability

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsFrameAcknowledgeCapabilitySet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdprfx/e4d498fd-822b-408d-b8b3-1c216f21265b
type TsFrameAcknowledgeCapabilitySet struct {
	MaxUnacknowledgedFrameCount uint32
}

func (c *TsFrameAcknowledgeCapabilitySet) Type() uint16 {
	return CAPSSETTYPE_FRAME_ACKNOWLEDGE
}

func (c *TsFrameAcknowledgeCapabilitySet) Read(r io.Reader) TsCapsSet {
	return core.ReadLE(r, c)
}

func (c *TsFrameAcknowledgeCapabilitySet) Write(w io.Writer) {
	core.WriteLE(w, c)
}
