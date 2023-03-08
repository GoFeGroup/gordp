package licPdu

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/sec"
	"io"
)

/* Client Info Packet Flags */
const (
	INFO_MOUSE                  = 0x00000001
	INFO_DISABLECTRLALTDEL      = 0x00000002
	INFO_AUTOLOGON              = 0x00000008
	INFO_UNICODE                = 0x00000010
	INFO_MAXIMIZESHELL          = 0x00000020
	INFO_LOGONNOTIFY            = 0x00000040
	INFO_COMPRESSION            = 0x00000080
	INFO_ENABLEWINDOWSKEY       = 0x00000100
	INFO_REMOTECONSOLEAUDIO     = 0x00002000
	INFO_FORCE_ENCRYPTED_CS_PDU = 0x00004000
	INFO_RAIL                   = 0x00008000
	INFO_LOGONERRORS            = 0x00010000
	INFO_MOUSE_HAS_WHEEL        = 0x00020000
	INFO_PASSWORD_IS_SC_PIN     = 0x00040000
	INFO_NOAUDIOPLAYBACK        = 0x00080000
	INFO_USING_SAVED_CREDS      = 0x00100000
	INFO_AUDIOCAPTURE           = 0x00200000
	INFO_VIDEO_DISABLE          = 0x00400000
	INFO_HIDEF_RAIL_SUPPORTED   = 0x02000000
)

// TsInfoPacket
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/732394f5-e2b5-4ac5-8a0a-35345386b0d1
type TsInfoPacket struct {
	CodePage         uint32
	Flag             uint32
	CbDomain         uint16
	CbUserName       uint16
	CbPassword       uint16
	CbAlternateShell uint16
	CbWorkingDir     uint16
	Domain           []byte
	UserName         []byte
	Password         []byte
	AlternateShell   []byte
	WorkingDir       []byte
	ExtendedInfo     *sec.TsExtendedInfoPacket
}

func (p *TsInfoPacket) Write(w io.Writer) {
	core.WriteLE(w, p.CodePage)
	core.WriteLE(w, p.Flag)
	core.WriteLE(w, p.CbDomain)
	core.WriteLE(w, p.CbUserName)
	core.WriteLE(w, p.CbPassword)
	core.WriteLE(w, p.CbAlternateShell)
	core.WriteLE(w, p.CbWorkingDir)
	core.WriteFull(w, p.Domain)
	core.WriteFull(w, p.UserName)
	core.WriteFull(w, p.Password)
	core.WriteFull(w, p.AlternateShell)
	core.WriteFull(w, p.WorkingDir)
	p.ExtendedInfo.Write(w)
}
