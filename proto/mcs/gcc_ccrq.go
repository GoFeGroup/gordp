package mcs

import (
	"bytes"
	"github.com/GoFeGroup/gordp/glog"
	"github.com/GoFeGroup/gordp/proto/mcs/per"
	"io"
)

var (
	t124_02_98_oid = []byte{0, 0, 20, 124, 0, 1}
	h221_cs_key    = "Duca"
	h221_sc_key    = "McDn"
)

// GccConferenceCreateRequest
// Copy From FreeRDP C.gcc_write_conference_create_request
// But I don't know what the CCR structure is.
type GccConferenceCreateRequest struct {
}

func (req *GccConferenceCreateRequest) Write(w io.Writer, userData []byte) {
	per.WriteChoice(w, 0)                        // From Key select object (0) of type OBJECT_IDENTIFIER
	per.WriteObjectIdentifier(w, t124_02_98_oid) // ITU-T T.124 (02/98) OBJECT_IDENTIFIER
	per.WriteLength(w, len(userData)+14)         // connectPDU length
	per.WriteChoice(w, 0)                        // From ConnectGCCPDU select conferenceCreateRequest (0) of type ConferenceCreateRequest
	per.WriteSelection(w, 0x08)                  // select optional userData from ConferenceCreateRequest
	per.WriteNumericString(w, "1", 1)            // ConferenceName::numeric
	per.WritePadding(w, 1)                       // padding
	per.WriteNumberOfSet(w, 1)                   // one set of UserData
	per.WriteChoice(w, 0xC0)                     // UserData::value present + select h221NonStandard (1)
	per.WriteOctetString(w, h221_cs_key, 4)      // UserData::value present + select h221NonStandard (1)
	per.WriteOctetString(w, string(userData), 0)
}

func (req *GccConferenceCreateRequest) Serialize(userData []byte) []byte {
	buff := new(bytes.Buffer)
	req.Write(buff, userData)
	glog.Debugf("ccr: len: %v, userData: %v", buff.Len(), len(userData))
	return buff.Bytes()
}
