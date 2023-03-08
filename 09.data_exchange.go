package gordp

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/t128"
)

func (c *Client) readPdu() t128.PDU {
	glog.Debugf("before peek")
	defer func() { glog.Debugf("exit readPDU") }()
	d := c.stream.Peek(1)
	switch d[0] {
	case 3:
		glog.Debugf("read tpkt pdu begin")
		return t128.ReadPDU(c.stream)
	case 0:
		glog.Debugf("read fastpath pdu begin")
		return t128.ReadFastPathPDU(c.stream)
	default:
		core.Throw("invalid package")
	}
	return nil
}
