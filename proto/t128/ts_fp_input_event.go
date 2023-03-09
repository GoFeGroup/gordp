package t128

/* FastPath Input Events */
const (
	FASTPATH_INPUT_EVENT_SCANCODE = 0x0
	FASTPATH_INPUT_EVENT_MOUSE    = 0x1
	FASTPATH_INPUT_EVENT_MOUSEX   = 0x2
	FASTPATH_INPUT_EVENT_SYNC     = 0x3
	FASTPATH_INPUT_EVENT_UNICODE  = 0x4
)

// TsFpInputEvent
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/76c4dd59-7ba0-445d-a03c-885212ab80f6
type TsFpInputEvent interface {
	Serialize() []byte
	iInputEvent()
}
