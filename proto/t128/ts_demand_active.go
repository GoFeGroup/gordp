package t128

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/capability"
	"io"
)

// TsDemandActivePduData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/bd612af5-cb54-43a2-9646-438bc3ecf5db
type TsDemandActivePduData struct {
	SharedId                   uint32
	LengthSourceDescriptor     uint16
	LengthCombinedCapabilities uint16
	SourceDescriptor           []byte
	NumberCapabilities         uint16
	Pad2Octets                 uint16
	CapabilitySets             []capability.TsCapsSet
	SessionId                  uint32
}

func (d *TsDemandActivePduData) Type() uint16 {
	return PDUTYPE_DEMANDACTIVEPDU
}

func (d *TsDemandActivePduData) iPDU() {}

func (d *TsDemandActivePduData) Serialize() []byte {
	core.Throw("not implement")
	return nil
}

func (d *TsDemandActivePduData) Read(r io.Reader) PDU {
	core.ReadLE(r, &d.SharedId)
	core.ReadLE(r, &d.LengthSourceDescriptor)
	core.ReadLE(r, &d.LengthCombinedCapabilities)
	d.SourceDescriptor = core.ReadBytes(r, int(d.LengthSourceDescriptor))
	core.ReadLE(r, &d.NumberCapabilities)
	core.ReadLE(r, &d.Pad2Octets)
	d.CapabilitySets = make([]capability.TsCapsSet, d.NumberCapabilities)
	for i := 0; i < int(d.NumberCapabilities); i++ {
		d.CapabilitySets = append(d.CapabilitySets, capability.Read(r))
	}
	core.ReadLE(r, &d.SessionId)
	return d
}
