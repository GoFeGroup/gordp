package t128

// TsFpKeyboardEvent
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/089d362b-31eb-4a1a-b6fa-92fe61bb5dbf
type TsFpKeyboardEvent struct {
}

func (e *TsFpKeyboardEvent) iInputEvent() {}

func (e *TsFpKeyboardEvent) Serialize() []byte {
	return nil
}
