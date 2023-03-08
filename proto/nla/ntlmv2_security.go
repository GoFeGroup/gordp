package nla

import (
	"bytes"
	"crypto/rc4"
	"github.com/GoFeGroup/gordp/core"
)

type NTLMv2Security struct {
	EncryptRC4 *rc4.Cipher
	DecryptRC4 *rc4.Cipher
	SigningKey []byte
	VerifyKey  []byte
	SeqNum     uint32
}

func (n *NTLMv2Security) Serialize(pubKey []byte) []byte {
	p := make([]byte, len(pubKey))
	n.EncryptRC4.XORKeyStream(p, pubKey)

	w := new(bytes.Buffer)
	core.WriteLE(w, n.SeqNum)
	core.WriteFull(w, pubKey)

	s1 := core.HMAC_MD5(n.SigningKey, w.Bytes())[:8]
	checksum := make([]byte, 8)
	n.EncryptRC4.XORKeyStream(checksum, s1)
	w.Reset()

	core.WriteLE(w, uint32(0x00000001))
	core.WriteFull(w, checksum)
	core.WriteLE(w, n.SeqNum)
	core.WriteFull(w, p)

	n.SeqNum++
	return w.Bytes()
}
