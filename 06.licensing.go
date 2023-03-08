package gordp

import "github.com/GoFeGroup/gordp/proto/pdu/licPdu"

func (c *Client) readLicensing() {
	licensing := licPdu.ServerLicensingPDU{}
	licensing.Read(c.stream)
}
