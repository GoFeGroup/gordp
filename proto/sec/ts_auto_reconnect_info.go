package sec

import (
	"github.com/GoFeGroup/gordp/core"
	"io"
)

// TsAutoReconnectInfo reconnect information
type TsAutoReconnectInfo struct{}

func (i *TsAutoReconnectInfo) Write(w io.Writer) {
	core.Throw("not implement") // FIXME: never used?
}
