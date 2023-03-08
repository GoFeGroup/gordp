package gordp

import (
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/t128"
)

func (c *Client) sendClientFinalization() {
	t128.WriteDataPdu(c.stream, c.userId, c.shareId, t128.NewTsSynchronizePduData(c.userId))
	t128.WriteDataPdu(c.stream, c.userId, c.shareId, &t128.TsControlPDU{Action: t128.CTRLACTION_COOPERATE})
	t128.WriteDataPdu(c.stream, c.userId, c.shareId, &t128.TsControlPDU{Action: t128.CTRLACTION_REQUEST_CONTROL})
	t128.WriteDataPdu(c.stream, c.userId, c.shareId, &t128.TsFontListPDU{ListFlags: 0x0003, EntrySize: 0x0032})

	t128.ReadExpectedDataPDU(c.stream, t128.PDUTYPE2_SYNCHRONIZE)
	ctlPdu1 := t128.ReadExpectedDataPDU(c.stream, t128.PDUTYPE2_CONTROL).(*t128.TsControlPDU)
	core.ThrowIf(ctlPdu1.Action != t128.CTRLACTION_COOPERATE, "server action is not `COOPERATE`")
	ctlPdu2 := t128.ReadExpectedDataPDU(c.stream, t128.PDUTYPE2_CONTROL).(*t128.TsControlPDU)
	core.ThrowIf(ctlPdu2.Action != t128.CTRLACTION_GRANTED_CONTROL, "server action is not `GRANTED_CONTROL`")
	t128.ReadExpectedDataPDU(c.stream, t128.PDUTYPE2_FONTMAP)
}
