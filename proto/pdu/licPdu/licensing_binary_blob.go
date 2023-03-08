package licPdu

// WBlobType
const (
	BB_ANY_BLOB                 = 0x0000
	BB_DATA_BLOB                = 0x0001
	BB_RANDOM_BLOB              = 0x0002
	BB_CERTIFICATE_BLOB         = 0x0003
	BB_ERROR_BLOB               = 0x0004
	BB_ENCRYPTED_DATA_BLOB      = 0x0009
	BB_KEY_EXCHG_ALG_BLOB       = 0x000D
	BB_SCOPE_BLOB               = 0x000E
	BB_CLIENT_USER_NAME_BLOB    = 0x000F
	BB_CLIENT_MACHINE_NAME_BLOB = 0x0010
)

// LicensingBinaryBlob
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/0a1c6079-af4d-4742-bb64-fecb8fa6e1a0
type LicensingBinaryBlob struct {
	WBlobType uint16
	WBlobLen  uint16
	BlobData  []byte
}
