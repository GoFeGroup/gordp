package gordp

import (
	"time"

	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/bitmap"
	"github.com/GoFeGroup/gordp/proto/t128"
)

type Option struct {
	Addr     string
	UserName string
	Password string

	ConnectTimeout time.Duration
}

type Processor interface {
	ProcessBitmap(*bitmap.Option, *bitmap.BitMap)
}

type Client struct {
	option Option

	//conn   net.Conn // TCP连接
	stream *core.Stream

	// from negotiation
	selectProtocol uint32 // 协商RDP协议，0:rdp, 1:ssl, 2:hybrid
	userId         uint16
	shareId        uint32
	serverVersion  uint32 // 服务端RDP版本号
}

func NewClient(opt *Option) *Client {
	c := &Client{
		option: Option{
			Addr:           opt.Addr,
			UserName:       opt.UserName,
			Password:       opt.Password,
			ConnectTimeout: opt.ConnectTimeout,
		},
	}
	if c.option.ConnectTimeout == 0 {
		c.option.ConnectTimeout = 5 * time.Second
	}
	return c
}

//func (c *Client) tcpConnect() {
//	conn, err := net.DialTimeout("tcp", c.option.Addr, c.option.ConnectTimeout)
//	core.ThrowError(err)
//	c.conn = conn
//}

// Connect
// https://www.cyberark.com/resources/threat-research-blog/explain-like-i-m-5-remote-desktop-protocol-rdp
func (c *Client) Connect() error {
	return core.Try(func() {
		c.stream = core.NewStream(c.option.Addr, c.option.ConnectTimeout)
		c.negotiation()
		c.basicSettingsExchange()
		c.channelConnect()
		c.sendClientInfo()
		c.readLicensing()
		c.capabilitiesExchange()
		c.sendClientFinalization()
	})
}

func (c *Client) Close() {
	c.stream.Close()
}

func (c *Client) Run(processor Processor) error {
	return core.Try(func() {
		for {
			pdu := c.readPdu()
			switch p := pdu.(type) {
			case *t128.TsFpUpdatePDU:
				if p.Length == 0 {
					break
				}
				switch pp := p.PDU.(type) {
				case *t128.TsFpUpdateBitmap:
					for _, v := range pp.Rectangles {
						option := &bitmap.Option{
							Top:         int(v.DestTop),  // for position
							Left:        int(v.DestLeft), // for position
							Width:       int(v.Width),
							Height:      int(v.Height),
							BitPerPixel: int(v.BitsPerPixel),
							Data:        v.BitmapDataStream,
						}
						// Note:
						// 1. 未压缩的位图数据被格式化为自底向上、从左到右的一系列像素。每个像素是字节的整数。每行包含四个字节的倍数（必要时最多包含三个字节的填充）
						// 2. 非32bpp格式的压缩位图使用交织RLE压缩并封装在RLE压缩位图流结构中
						//    ->  https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/b3b60873-16a8-4cbc-8aaa-5f0a93083280
						// 3. 颜色深度为32bpp的压缩位图则使用RDP 6.0位图压缩压缩并存储在RDP 6.0 Bitmap压缩流结构内
						//    ->  https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpegdi/9b422f69-8e05-4c6d-b6fb-fa02ef75a8f2

						if v.BitsPerPixel == 32 {
							processor.ProcessBitmap(option, bitmap.NewBitMapFromRDP6(option))
						} else {
							processor.ProcessBitmap(option, bitmap.NewBitmapFromRLE(option))
						}
					}
				default:
					glog.Debugf("pdutype2: %T", pp)
				}
			default:
				glog.Debugf("type: %T", p)
			}
		}
	})
}
