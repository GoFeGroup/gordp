package capability

import (
	"bytes"
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

const (
	CAPSTYPE_GENERAL                 = 0x0001
	CAPSTYPE_BITMAP                  = 0x0002
	CAPSTYPE_ORDER                   = 0x0003
	CAPSTYPE_BITMAPCACHE             = 0x0004
	CAPSTYPE_CONTROL                 = 0x0005
	CAPSTYPE_ACTIVATION              = 0x0007
	CAPSTYPE_POINTER                 = 0x0008
	CAPSTYPE_SHARE                   = 0x0009
	CAPSTYPE_COLORCACHE              = 0x000A
	CAPSTYPE_SOUND                   = 0x000C
	CAPSTYPE_INPUT                   = 0x000D
	CAPSTYPE_FONT                    = 0x000E
	CAPSTYPE_BRUSH                   = 0x000F
	CAPSTYPE_GLYPHCACHE              = 0x0010
	CAPSTYPE_OFFSCREENCACHE          = 0x0011
	CAPSTYPE_BITMAPCACHE_HOSTSUPPORT = 0x0012
	CAPSTYPE_BITMAPCACHE_REV2        = 0x0013
	CAPSTYPE_VIRTUALCHANNEL          = 0x0014
	CAPSTYPE_DRAWNINEGRIDCACHE       = 0x0015
	CAPSTYPE_DRAWGDIPLUS             = 0x0016
	CAPSTYPE_RAIL                    = 0x0017
	CAPSTYPE_WINDOW                  = 0x0018
	CAPSETTYPE_COMPDESK              = 0x0019
	CAPSETTYPE_MULTIFRAGMENTUPDATE   = 0x001A
	CAPSETTYPE_LARGE_POINTER         = 0x001B
	CAPSETTYPE_SURFACE_COMMANDS      = 0x001C
	CAPSETTYPE_BITMAP_CODECS         = 0x001D
	CAPSSETTYPE_FRAME_ACKNOWLEDGE    = 0x001E
)

var capsMap = map[uint16]TsCapsSet{
	CAPSTYPE_GENERAL:                 &TsGeneralCapabilitySet{},
	CAPSTYPE_BITMAP:                  &TsBitmapCapabilitySet{},
	CAPSTYPE_ORDER:                   &TsOrderCapabilitySet{},
	CAPSTYPE_BITMAPCACHE:             &TsBitmapCacheCapabilitySet{},
	CAPSTYPE_CONTROL:                 &TsControlCapabilitySet{},
	CAPSTYPE_ACTIVATION:              &TsWindowActivationCapabilitySet{},
	CAPSTYPE_POINTER:                 &TsPointerCapabilitySet{},
	CAPSTYPE_SHARE:                   &TsShareCapabilitySet{},
	CAPSTYPE_COLORCACHE:              &TsColorTableCapabilitySet{},
	CAPSTYPE_SOUND:                   &TsSoundCapabilitySet{},
	CAPSTYPE_INPUT:                   &TsInputCapabilitySet{},
	CAPSTYPE_FONT:                    &TsFontCapabilitySet{},
	CAPSTYPE_BRUSH:                   &TsBrushCapabilitySet{},
	CAPSTYPE_GLYPHCACHE:              &TsGlyphCacheCapabilitySet{},
	CAPSTYPE_OFFSCREENCACHE:          &TsOffscreenCapabilitySet{},
	CAPSTYPE_BITMAPCACHE_HOSTSUPPORT: &TsBitmapCacheHostSupportCapabilitySet{},
	CAPSTYPE_BITMAPCACHE_REV2:        &TsBitmapCacheCapabilitySetRev2{},
	CAPSTYPE_VIRTUALCHANNEL:          &TsVirtualChannelCapabilitySet{},
	CAPSTYPE_DRAWNINEGRIDCACHE:       &TsDrawNineGridCapabilitySet{},
	CAPSTYPE_DRAWGDIPLUS:             &TsDrawGdiPlusCapabilitySet{},
	CAPSTYPE_RAIL:                    &RemoteProgramsCapabilitySet{},
	CAPSTYPE_WINDOW:                  &WindowListCapabilitySet{},
	CAPSETTYPE_COMPDESK:              &TsCompDeskCapabilitySet{},
	CAPSETTYPE_MULTIFRAGMENTUPDATE:   &TsMultiFragmentUpdateCapabilitySet{},
	CAPSETTYPE_LARGE_POINTER:         &TsLargePointerCapabilitySet{},
	CAPSETTYPE_SURFACE_COMMANDS:      &TsSurfCmdsCapabilitySet{},
	CAPSETTYPE_BITMAP_CODECS:         &TsBitmapCodecsCapabilitySet{},
	CAPSSETTYPE_FRAME_ACKNOWLEDGE:    &TsFrameAcknowledgeCapabilitySet{},
}

type TsCapsSetHeader struct {
	CapabilitySetType uint16
	LengthCapability  uint16
}

func (h *TsCapsSetHeader) Read(r io.Reader) {
	core.ReadLE(r, h)
}

// TsCapsSet
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/d705c3b6-a392-4b32-9610-391f6af62323
type TsCapsSet interface {
	Type() uint16
	Read(r io.Reader) TsCapsSet
	Write(w io.Writer)
}

func Read(r io.Reader) TsCapsSet {
	header := &TsCapsSetHeader{}
	header.Read(r)

	data := make([]byte, header.LengthCapability-4)
	_, err := io.ReadFull(r, data)
	core.ThrowError(err)
	r = bytes.NewReader(data)

	glog.Debugf("capability type: %0#4x", header.CapabilitySetType)
	c, ok := capsMap[header.CapabilitySetType]
	core.ThrowIf(!ok, fmt.Errorf("capability type: %0#4x, need implement", header.CapabilitySetType))
	return c.Read(r)
}

func Write(w io.Writer, caps []TsCapsSet) {
	for _, c := range caps {
		core.WriteLE(w, c.Type())
		capBytes := capSerialize(c)
		glog.Debugf("cap: %x, %v, %x", c.Type(), len(capBytes), capBytes)
		core.WriteLE(w, uint16(len(capBytes)+4))
		core.WriteFull(w, capBytes)
	}
}

func capSerialize(c TsCapsSet) []byte {
	buff := new(bytes.Buffer)
	c.Write(buff)
	return buff.Bytes()
}

func Serialize(caps []TsCapsSet) []byte {
	buff := new(bytes.Buffer)
	Write(buff, caps)
	return buff.Bytes()
}
