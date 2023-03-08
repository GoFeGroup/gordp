package gordp

import (
	"github.com/GoFeGroup/gordp/proto/t128"
)

func (c *Client) capabilitiesExchange() {
	demandActivePDU := t128.ReadExpectedPDU(c.stream, t128.PDUTYPE_DEMANDACTIVEPDU).(*t128.TsDemandActivePduData)
	confirmActivePduData := t128.NewTsConfirmActivePduData(demandActivePDU)
	c.shareId = demandActivePDU.SharedId
	t128.WritePDU(c.stream, c.userId, confirmActivePduData)
}
