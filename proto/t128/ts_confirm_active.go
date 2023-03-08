package t128

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/capability"
	"io"
)

// TsConfirmActivePduData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/4e9722c3-ad83-43f5-af5a-529f73d88b48
type TsConfirmActivePduData struct {
	SharedId                   uint32
	OriginatorId               uint16 //This field MUST be set to the server channel ID (0x03EA).
	LengthSourceDescriptor     uint16
	LengthCombinedCapabilities uint16
	SourceDescriptor           []byte //A variable-length array of bytes containing a source descriptor
	NumberCapabilities         uint16
	Pad2Octets                 uint16
	CapabilitySets             []capability.TsCapsSet
}

func (d *TsConfirmActivePduData) Type() uint16 {
	return PDUTYPE_CONFIRMACTIVEPDU
}

func (d *TsConfirmActivePduData) iPDU() {}

func (d *TsConfirmActivePduData) Read(r io.Reader) PDU {
	core.Throw("not implement")
	return d
}

func (d *TsConfirmActivePduData) Write(w io.Writer) {
	capsBytes := capability.Serialize(d.CapabilitySets)
	d.LengthCombinedCapabilities = uint16(len(capsBytes)) + 2 + 2 // FIXME
	d.NumberCapabilities = uint16(len(d.CapabilitySets))
	core.WriteLE(w, d.SharedId)
	core.WriteLE(w, d.OriginatorId)
	core.WriteLE(w, d.LengthSourceDescriptor)
	core.WriteLE(w, d.LengthCombinedCapabilities)
	core.WriteFull(w, d.SourceDescriptor)
	core.WriteLE(w, d.NumberCapabilities)
	core.WriteLE(w, d.Pad2Octets)
	core.WriteFull(w, capsBytes)

	glog.Debugf("caps: %v", len(d.CapabilitySets))
	glog.Debugf("capsByte: %v - %x", len(capsBytes), capsBytes)
}

func (d *TsConfirmActivePduData) Serialize() []byte {
	buff := new(bytes.Buffer)
	d.Write(buff)
	return buff.Bytes()
}

func NewTsConfirmActivePduData(demandActivePdu *TsDemandActivePduData) *TsConfirmActivePduData {
	confirmActiveData := &TsConfirmActivePduData{
		OriginatorId:               0x03EA,
		SharedId:                   demandActivePdu.SharedId,
		LengthSourceDescriptor:     demandActivePdu.LengthSourceDescriptor,
		SourceDescriptor:           demandActivePdu.SourceDescriptor,
		LengthCombinedCapabilities: demandActivePdu.LengthCombinedCapabilities,
		CapabilitySets: []capability.TsCapsSet{
			capability.NewTsGeneralCapabilitySet(),
			capability.NewTsBitmapCapabilitySet(),
			capability.NewTsOrderCapabilitySet(),
			&capability.TsBitmapCacheCapabilitySet{},
			&capability.TsPointerCapabilitySet{ColorPointerCacheSize: 20},
			capability.NewTsInputCapabilitySet(),
			&capability.TsBrushCapabilitySet{},
			&capability.TsGlyphCacheCapabilitySet{},
			&capability.TsOffscreenCapabilitySet{},
			&capability.TsVirtualChannelCapabilitySet{},
			&capability.TsSoundCapabilitySet{},
			&capability.TsMultiFragmentUpdateCapabilitySet{},
			capability.NewRemoteProgramsCapabilitySet(),
		},
	}

	return confirmActiveData
}
