package gordp

import (
	"fmt"
	"os"
	"testing"

	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/bitmap"
)

type processor struct {
	i int
}

func (p *processor) ProcessBitmap(option *bitmap.Option, bitmap *bitmap.BitMap) {
	p.i++
	_ = os.MkdirAll("./png", 0755)
	_ = os.WriteFile(fmt.Sprintf("./png/%v.png", p.i), bitmap.ToPng(), 0644)
}

func TestRdpConnect(t *testing.T) {
	client := NewClient(&Option{
		Addr:     "10.226.239.200:3389",
		UserName: "administrator",
		Password: "[YourPasswordHere]",
	})
	core.ThrowError(client.Connect())
	core.ThrowError(client.Run(&processor{}))
}
