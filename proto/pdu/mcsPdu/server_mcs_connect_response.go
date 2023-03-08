package mcsPdu

import (
	"bytes"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/x224"
	"io"
)

// ServerMcsConnectResponsePDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/927de44c-7fe8-4206-a14f-e5517dc24b1c
type ServerMcsConnectResponsePDU struct {
	McsCrsp                         mcs.ConnectResponse
	GccCrsp                         mcs.GccConferenceCreateResponse
	ServerCoreData                  mcs.ServerCoreData
	ServerNetworkData               mcs.ServerNetworkData
	ServerSecurityData              mcs.ServerSecurityData
	ServerMessageChannelData        mcs.ServerMessageChannelData
	ServerMultitransportChannelData mcs.ServerMultitransportChannelData
}

func (pdu *ServerMcsConnectResponsePDU) Read(r io.Reader) {
	data := x224.Read(r)
	glog.Debugf("recv McsConnectResponse: %v", len(data))
	pdu.McsCrsp.Load(data)
	glog.Debugf("McsConnectResponse userData: %v - %x", len(pdu.McsCrsp.UserData), pdu.McsCrsp.UserData)
	rd := bytes.NewReader(pdu.McsCrsp.UserData)
	data = pdu.GccCrsp.Read(rd)
	glog.Debugf("McsCRSP: userData: %v - %x", len(data), data)

	rd = bytes.NewReader(data)
	for {
		header := mcs.UserDataHeader{}
		if rd.Len() <= 0 {
			break
		}
		switch header.Read(rd); header.Type {
		case mcs.SC_CORE:
			pdu.ServerCoreData.Read(rd)
		case mcs.SC_SECURITY:
			pdu.ServerSecurityData.Read(rd)
		case mcs.SC_NET:
			pdu.ServerNetworkData.Read(rd)
		}
	}
}
