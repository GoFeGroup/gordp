package t128

import (
	"bytes"

	"github.com/GoFeGroup/gordp/core"
)

const (
	// Mouse wheel event:
	PTRFLAGS_HWHEEL         = 0x0400
	PTRFLAGS_WHEEL          = 0x0200
	PTRFLAGS_WHEEL_NEGATIVE = 0x0100
	WheelRotationMask       = 0x01FF

	// Mouse movement event:
	PTRFLAGS_MOVE = 0x0800

	// Mouse button events:
	PTRFLAGS_DOWN    = 0x8000
	PTRFLAGS_BUTTON1 = 0x1000
	PTRFLAGS_BUTTON2 = 0x2000
	PTRFLAGS_BUTTON3 = 0x4000
)

// TsFpPointerEvent
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/16a96ded-b3d3-4468-b993-9c7a51297510
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/2c1ced34-340a-46cd-be6e-fc8cab7c3b17
type TsFpPointerEvent struct {
	PointerFlags uint16
	XPos, YPos   uint16
}

func (e *TsFpPointerEvent) iInputEvent() {}

func (e *TsFpPointerEvent) Serialize() []byte {
	buff := new(bytes.Buffer)
	core.WriteLE(buff, uint8(FASTPATH_INPUT_EVENT_MOUSE)) // eventHeader, eventFlags=0, eventCode=FASTPATH_INPUT_EVENT_MOUSE
	core.WriteLE(buff, e)
	return buff.Bytes()
}
