package nla

// NVersion
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/b1a6ceb2-f8ad-462b-b5af-f18527c48175
type NVersion struct {
	ProductMajorVersion uint8
	ProductMinorVersion uint8
	ProductBuild        uint16
	Reserved            [3]byte
	NTLMRevisionCurrent uint8
}

func NewNVersion() NVersion {
	return NVersion{
		ProductMajorVersion: WINDOWS_MAJOR_VERSION_6,
		ProductMinorVersion: WINDOWS_MINOR_VERSION_0,
		ProductBuild:        6002,
		NTLMRevisionCurrent: NTLMSSP_REVISION_W2K3,
	}
}
