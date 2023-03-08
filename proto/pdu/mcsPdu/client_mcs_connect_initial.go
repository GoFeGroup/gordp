package mcsPdu

import (
	"bytes"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs"
	"github.com/GoFeGroup/gordp/proto/x224"
	"io"
)

// ClientMcsConnectInitialPDU
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/db6713ee-1c0e-4064-a3b3-0fac30b4037b
type ClientMcsConnectInitialPDU struct {
	McsCi                           *mcs.ConnectInitial            //  [T125] section 11.1
	GccCCrq                         mcs.GccConferenceCreateRequest //  [T124] section 8.7
	ClientCoreData                  *mcs.ClientCoreData
	ClientSecurityData              *mcs.ClientSecurityData
	ClientNetworkData               *mcs.ClientNetworkData
	ClientClusterData               interface{}
	ClientMonitorData               interface{}
	ClientMessageChannelData        interface{}
	ClientMultitransportChannelData interface{}
	ClientMonitorExtendedData       interface{}
}

func (pdu *ClientMcsConnectInitialPDU) Write(w io.Writer) {
	var arr [][]byte
	arr = append(arr, pdu.ClientCoreData.Serialize())
	arr = append(arr, pdu.ClientNetworkData.Serialize())
	arr = append(arr, pdu.ClientSecurityData.Serialize())
	pdu.McsCi.UserData = pdu.GccCCrq.Serialize(bytes.Join(arr, nil))
	glog.Debugf("GccCCrq: %x", pdu.McsCi.UserData)
	x224.Write(w, pdu.McsCi.Serialize())
}

func NewClientMcsConnectInitialPdu(selectedProtocol uint32) *ClientMcsConnectInitialPDU {
	pdu := &ClientMcsConnectInitialPDU{}
	pdu.McsCi = mcs.NewClientInitial()
	pdu.ClientCoreData = mcs.NewClientCoreData()
	pdu.ClientSecurityData = mcs.NewClientSecurityData()
	pdu.ClientNetworkData = mcs.NewClientNetworkData()
	pdu.ClientCoreData.ServerSelectedProtocol = selectedProtocol
	return pdu
}
