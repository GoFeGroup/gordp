package gordp

import (
	"fmt"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/bitmap"
	"image/png"
	"os"
	"testing"
)

type processor struct {
	i int
}

func (p *processor) ProcessBitmap(option *bitmap.Option, bitmap *bitmap.BitMap) {
	p.i++
	_ = os.MkdirAll("./png", 0755)
	file, err := os.Create(fmt.Sprintf("./png/%v.png", p.i))
	core.ThrowError(err)
	core.ThrowError(png.Encode(file, bitmap.Image))
	core.ThrowError(file.Close())
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
