package mcs

import (
	"bytes"
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/mcs/per"
	"io"
)

// GccConferenceCreateResponse
// Copy From FreeRDP C.gcc_read_conference_create_response
// But I don't know what the CCR structure is.
type GccConferenceCreateResponse struct {
}

func (res *GccConferenceCreateResponse) Read(r io.Reader) []byte {
	_ = per.ReadChoice(r)
	if oid := per.ReadObjectIdentifier(r); !bytes.Equal(oid, t124_02_98_oid) {
		core.Throw(fmt.Errorf("invalid oid: %x", oid))
	}
	_ = per.ReadLength(r)                             // ConnectData::connectPDU (OCTET_STRING)
	_ = per.ReadChoice(r)                             // ConnectGCCPDU
	_ = per.ReadInteger16(r, MCS_CHANNEL_USERID_BASE) //ConferenceCreateResponse::nodeID (UserID)
	_ = per.ReadInteger(r)                            //ConferenceCreateResponse::tag (INTEGER)
	enum := per.ReadEnumerated(r)                     //ConferenceCreateResponse::result (ENUMERATED)
	core.ThrowIf(enum+1 > 16, fmt.Errorf("PER invalid data, expect %0#x < %0#x", enum, 16))
	_ = per.ReadNumberOfSet(r) // number of UserData sets
	_ = per.ReadChoice(r)      // UserData::value present + select h221NonStandard (1)
	oStr := per.ReadOctetString(r, 4)
	core.ThrowIf(!bytes.Equal(oStr, []byte(h221_sc_key)), "invalid data")
	return per.ReadOctetString(r, 0) // userData (OCTET_STRING)
}
