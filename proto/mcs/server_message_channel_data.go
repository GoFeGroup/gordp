package mcs

// ServerMessageChannelData
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/9269d58a-3d85-48a2-942a-bb0bbe5a55aa
type ServerMessageChannelData struct {
	Header    uint32
	ChannelId uint16
}
