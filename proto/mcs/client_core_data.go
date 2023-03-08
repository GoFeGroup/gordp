package mcs

import (
	"github.com/GoFeGroup/gordp/core"
	"os"
)

// ClientCoreData Version
const (
	RDP_VERSION_4      = 0x00080001 // RDP 4.0 clients
	RDP_VERSION_5_PLUS = 0x00080004 // RDP 5.0, 5.1, 5.2, 6.0, 6.1, 7.0, 7.1, 8.0, and 8.1 clients
	RDP_VERSION_10     = 0x00080005 // RDP 10.0 clients
	RDP_VERSION_10_0   = 0x00080005
	RDP_VERSION_10_1   = 0x00080006
	RDP_VERSION_10_2   = 0x00080007
	RDP_VERSION_10_3   = 0x00080008
	RDP_VERSION_10_4   = 0x00080009
	RDP_VERSION_10_5   = 0x0008000a
	RDP_VERSION_10_6   = 0x0008000b
	RDP_VERSION_10_7   = 0x0008000C
	RDP_VERSION_10_8   = 0x0008000D
	RDP_VERSION_10_9   = 0x0008000E
	RDP_VERSION_10_10  = 0x0008000F
	RDP_VERSION_10_11  = 0x00080010
)
const (
	RNS_UD_COLOR_8BPP      = 0xCA01
	RNS_UD_COLOR_16BPP_555 = 0xCA02
	RNS_UD_COLOR_16BPP_565 = 0xCA03
	RNS_UD_COLOR_24BPP     = 0xCA04
)
const (
	RNS_UD_SAS_DEL = 0xAA03
)
const (
	ARABIC              uint32 = 0x00000401
	BULGARIAN                  = 0x00000402
	CHINESE_US_KEYBOARD        = 0x00000404
	CZECH                      = 0x00000405
	DANISH                     = 0x00000406
	GERMAN                     = 0x00000407
	GREEK                      = 0x00000408
	US                         = 0x00000409
	SPANISH                    = 0x0000040a
	FINNISH                    = 0x0000040b
	FRENCH                     = 0x0000040c
	HEBREW                     = 0x0000040d
	HUNGARIAN                  = 0x0000040e
	ICELANDIC                  = 0x0000040f
	ITALIAN                    = 0x00000410
	JAPANESE                   = 0x00000411
	KOREAN                     = 0x00000412
	DUTCH                      = 0x00000413
	NORWEGIAN                  = 0x00000414
)
const (
	KT_IBM_PC_XT_83_KEY = 0x00000001
	KT_OLIVETTI         = 0x00000002
	KT_IBM_PC_AT_84_KEY = 0x00000003
	KT_IBM_101_102_KEYS = 0x00000004
	KT_NOKIA_1050       = 0x00000005
	KT_NOKIA_9140       = 0x00000006
	KT_JAPANESE         = 0x00000007
)

const (
	HIGH_COLOR_4BPP  = 0x0004
	HIGH_COLOR_8BPP  = 0x0008
	HIGH_COLOR_15BPP = 0x000f
	HIGH_COLOR_16BPP = 0x0010
	HIGH_COLOR_24BPP = 0x0018
)

const (
	RNS_UD_24BPP_SUPPORT uint16 = 0x0001
	RNS_UD_16BPP_SUPPORT        = 0x0002
	RNS_UD_15BPP_SUPPORT        = 0x0004
	RNS_UD_32BPP_SUPPORT        = 0x0008
)

const (
	RNS_UD_CS_SUPPORT_ERRINFO_PDU        uint16 = 0x0001
	RNS_UD_CS_WANT_32BPP_SESSION                = 0x0002
	RNS_UD_CS_SUPPORT_STATUSINFO_PDU            = 0x0004
	RNS_UD_CS_STRONG_ASYMMETRIC_KEYS            = 0x0008
	RNS_UD_CS_UNUSED                            = 0x0010
	RNS_UD_CS_VALID_CONNECTION_TYPE             = 0x0020
	RNS_UD_CS_SUPPORT_MONITOR_LAYOUT_PDU        = 0x0040
	RNS_UD_CS_SUPPORT_NETCHAR_AUTODETECT        = 0x0080
	RNS_UD_CS_SUPPORT_DYNVC_GFX_PROTOCOL        = 0x0100
	RNS_UD_CS_SUPPORT_DYNAMIC_TIME_ZONE         = 0x0200
	RNS_UD_CS_SUPPORT_HEARTBEAT_PDU             = 0x0400
)

// ClientCoreData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/00f1da4a-ee9c-421a-852f-c19f92343d73
type ClientCoreData struct {
	Header                 UserDataHeader // CS_CORE
	Version                uint32         // RDP_VERSION_5_PLUS
	DesktopWidth           uint16
	DesktopHeight          uint16
	ColorDepth             uint16
	SasSequence            uint16
	KbdLayout              uint32
	ClientBuild            uint32
	ClientName             [32]byte
	KeyboardType           uint32
	KeyboardSubType        uint32
	KeyboardFnKeys         uint32
	ImeFileName            [64]byte
	PostBeta2ColorDepth    uint16
	ClientProductId        uint16
	SerialNumber           uint32
	HighColorDepth         uint16
	SupportedColorDepths   uint16
	EarlyCapabilityFlags   uint16
	ClientDigProductId     [64]byte
	ConnectionType         uint8
	Pad1octet              uint8
	ServerSelectedProtocol uint32
}

func (coreData *ClientCoreData) Serialize() []byte {
	return core.ToLE(coreData)
}

func NewClientCoreData() *ClientCoreData {
	name, _ := os.Hostname()
	coreData := &ClientCoreData{
		Header:                 UserDataHeader{Type: CS_CORE, Len: 0xd8},
		Version:                RDP_VERSION_5_PLUS,
		DesktopWidth:           1280,
		DesktopHeight:          800,
		ColorDepth:             RNS_UD_COLOR_8BPP,
		SasSequence:            RNS_UD_SAS_DEL,
		KbdLayout:              US,
		ClientBuild:            3790,
		ClientName:             [32]byte{},
		KeyboardType:           KT_IBM_101_102_KEYS,
		KeyboardSubType:        0,
		KeyboardFnKeys:         12,
		ImeFileName:            [64]byte{},
		PostBeta2ColorDepth:    RNS_UD_COLOR_8BPP,
		ClientProductId:        1,
		SerialNumber:           0,
		HighColorDepth:         HIGH_COLOR_24BPP,
		SupportedColorDepths:   RNS_UD_15BPP_SUPPORT | RNS_UD_16BPP_SUPPORT | RNS_UD_24BPP_SUPPORT | RNS_UD_32BPP_SUPPORT,
		EarlyCapabilityFlags:   RNS_UD_CS_SUPPORT_ERRINFO_PDU,
		ClientDigProductId:     [64]byte{},
		ConnectionType:         0,
		Pad1octet:              0,
		ServerSelectedProtocol: 0,
	}
	copy(coreData.ClientName[:], core.UnicodeEncode(name))
	return coreData
}
