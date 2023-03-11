package mcs

import (
	"io"

	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/glog"
)

// RSAPublicKey
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/fe93545c-772a-4ade-9d02-ad1e0d81b6af
type RSAPublicKey struct {
	Magic   uint32 // 0x31415352
	KeyLen  uint32 // MUST be ((BitLen / 8) + 8) bytes.
	BitLen  uint32 // The number of bits in the public key modulus.
	DataLen uint32 // This value is directly related to the BitLen field and MUST be ((BitLen / 8) - 1) bytes.
	PubExp  uint32 // The public exponent of the public key.
	Modulus []byte // The modulus field contains all (BitLen / 8) bytes of the public key
	Padding []byte // 8字节对齐
}

// ProprietaryServerCertificate
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/a37d449a-73ac-4f00-9b9d-56cefc954634
type ProprietaryServerCertificate struct {
	DwSigAlgId        uint32 // This field MUST be set to SIGNATURE_ALG_RSA (0x00000001).
	DwKeyAlgId        uint32 // This field MUST be set to KEY_EXCHANGE_ALG_RSA (0x00000001).
	PublicKeyBlobType uint16 // This field MUST be set to BB_RSA_KEY_BLOB (0x0006).
	PublicKeyBlobLen  uint16 // The size in bytes of the PublicKeyBlob field.
	PublicKeyBlob     RSAPublicKey
	SignatureBlobType uint16 // This field is set to BB_RSA_SIGNATURE_BLOB (0x0008).
	SignatureBlobLen  uint16 // The size in bytes of the SignatureBlob field.
	SignatureBlob     []byte
	Padding           []byte
}

func (p *ProprietaryServerCertificate) GetPublicKey() (uint32, []byte) {
	return p.PublicKeyBlob.PubExp, p.PublicKeyBlob.Modulus
}
func (p *ProprietaryServerCertificate) Verify() bool {
	return true // TODO:
}

func (p *ProprietaryServerCertificate) Read(r io.Reader) {
	core.ReadLE(r, &p.DwSigAlgId)        // 1
	core.ReadLE(r, &p.DwKeyAlgId)        // 1
	core.ReadLE(r, &p.PublicKeyBlobType) // 6
	core.ReadLE(r, &p.PublicKeyBlobLen)  // 92

	core.ReadLE(r, &p.PublicKeyBlob.Magic)   // 826364754
	core.ReadLE(r, &p.PublicKeyBlob.KeyLen)  // 72 (BitLen/8+8)
	core.ReadLE(r, &p.PublicKeyBlob.BitLen)  // 512
	core.ReadLE(r, &p.PublicKeyBlob.DataLen) // 63 (BitLen/8-1)
	core.ReadLE(r, &p.PublicKeyBlob.PubExp)  // 65537
	p.PublicKeyBlob.Modulus = core.ReadBytes(r, int(p.PublicKeyBlob.KeyLen))
	//core.ReadLE(r, &p.PublicKeyBlob.Modulus)
	//core.ReadLE(r, &p.PublicKeyBlob.Padding)
	//p.PublicKeyBlob.Padding = core.ReadBytes(r, int(8-p.PublicKeyBlob.KeyLen%8))

	core.ReadLE(r, &p.SignatureBlobType) // 0x0008
	core.ReadLE(r, &p.SignatureBlobLen)  // 72
	glog.Debugf("%+v", p)
	p.SignatureBlob = core.ReadBytes(r, int(p.SignatureBlobLen))
	//p.Padding = core.ReadBytes(r, int(8-p.SignatureBlobLen%8))

	//core.ReadLE(r, &p.SignatureBlob)
	//core.ReadLE(r, &p.Padding)
}
