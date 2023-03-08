package mcs

import (
	"bytes"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"io"
)

// ChannelDef Options
const (
	CHANNEL_OPTION_INITIALIZED   = 0x80000000
	CHANNEL_OPTION_ENCRYPT_RDP   = 0x40000000
	CHANNEL_OPTION_ENCRYPT_SC    = 0x20000000
	CHANNEL_OPTION_ENCRYPT_CS    = 0x10000000
	CHANNEL_OPTION_PRI_HIGH      = 0x08000000
	CHANNEL_OPTION_PRI_MED       = 0x04000000
	CHANNEL_OPTION_PRI_LOW       = 0x02000000
	CHANNEL_OPTION_COMPRESS_RDP  = 0x00800000
	CHANNEL_OPTION_COMPRESS      = 0x00400000
	CHANNEL_OPTION_SHOW_PROTOCOL = 0x00200000
	REMOTE_CONTROL_PERSISTENT    = 0x00100000
)

// ChannelDef
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/a9b9dc4a-6ae4-4b04-a6f5-87e5ed6fd7e7
type ChannelDef struct {
	Name    [8]byte
	Options uint32
}

func (d *ChannelDef) Write(w io.Writer) {
	core.Throw("not implement")
}

// ClientNetworkData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/49f99e00-caf1-4786-b43c-d425de29a03f
type ClientNetworkData struct {
	Header          UserDataHeader // 	CS_NET      = 0xC003
	ChannelCount    uint32
	ChannelDefArray []ChannelDef
}

func (networkData *ClientNetworkData) Serialize() []byte {
	buff := new(bytes.Buffer)
	core.WriteLE(buff, networkData.Header)
	core.WriteLE(buff, networkData.ChannelCount)
	for _, v := range networkData.ChannelDefArray {
		v.Write(buff)
	}
	glog.Debugf("clientNetworkData: %x", buff.Bytes())
	return buff.Bytes()
}

func NewClientNetworkData() *ClientNetworkData {
	return &ClientNetworkData{
		Header: UserDataHeader{Type: CS_NET, Len: 0x08},
	}
}
