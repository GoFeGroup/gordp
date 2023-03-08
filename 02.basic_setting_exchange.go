package gordp

import (
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/pdu/mcsPdu"
)

func (c *Client) basicSettingsExchange() {
	mcsReqPdu := mcsPdu.NewClientMcsConnectInitialPdu(c.selectProtocol)
	mcsReqPdu.Write(c.stream)
	glog.Debugf("send connect initial pdu ok.")

	mcsResPdu := mcsPdu.ServerMcsConnectResponsePDU{}
	mcsResPdu.Read(c.stream)
	glog.Debugf("receive connect response pdu ok")
	glog.Debugf("rdp version: client=%0#x, server=%0#x", mcsReqPdu.ClientCoreData.Version, mcsResPdu.ServerCoreData.Version)
	c.serverVersion = mcsResPdu.ServerCoreData.Version
}
