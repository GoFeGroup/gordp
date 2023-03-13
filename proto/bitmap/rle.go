package bitmap

import (
	"bytes"
	"image"
	"image/color"
	"io"

	"github.com/GoFeGroup/gordp/core"
)

const (
	g_MaskRegularRunLength = 0x1F
	g_MaskLiteRunLength    = 0x0F
	g_MaskSpecialFgBg1     = 0x03
	g_MaskSpecialFgBg2     = 0x05

	// opcode>>=5,count=opcode&0x1F:input[0]+32
	REGULAR_BG_RUN      = 0x00 // opcode=FILL
	REGULAR_FG_RUN      = 0x01 // opcode=MIX
	REGULAR_FGBG_IMAGE  = 0x02 // count=(code&0xF)<<3:input[0]+1,opcode=FOM
	REGULAR_COLOR_RUN   = 0x03 // opcode=COLOR
	REGULAR_COLOR_IMAGE = 0x04 // opcode=COPY

	// >> shift >> 4
	LITE_SET_FG_FG_RUN     = 0x0C // count=code&0xF:count=input[0]+16,opcode=MIX_SET
	LITE_SET_FG_FGBG_IMAGE = 0x0D // count=(code&0xF)<<3:input[0]+1,opcode=FOM_SET
	LITE_DITHERED_RUN      = 0x0E // count=code&0xF:count=input[0]+16,opcode=BICOLOR

	// no >> shift
	MEGA_MEGA_BG_RUN         = 0xF0 // count=input[0]|(input[1]<<8),opcode=FILL
	MEGA_MEGA_FG_RUN         = 0xF1 // count=input[0]|(input[1]<<8),opcode=MIX
	MEGA_MEGA_FGBG_IMAGE     = 0xF2 // count=input[0]|(input[1]<<8),opcode=FOM
	MEGA_MEGA_COLOR_RUN      = 0xF3 // count=input[0]|(input[1]<<8),opcode=COLOR
	MEGA_MEGA_COLOR_IMAGE    = 0xF4 // count=input[0]|(input[1]<<8),opcode=COPY
	MEGA_MEGA_SET_FG_RUN     = 0xF6 // count=input[0]|(input[1]<<8),opcode=MIX_SET
	MEGA_MEGA_SET_FGBG_IMAGE = 0xF7 // count=input[0]|(input[1]<<8),opcode=FOM_SET
	MEGA_MEGA_DITHERED_RUN   = 0xF8 // count=input[0]|(input[1]<<8),opcode=BICOLOR,input=2
	SPECIAL_FGBG_1           = 0xF9 // count=8,opcode=SPECIAL_FGBG_1
	SPECIAL_FGBG_2           = 0xFA // count=8,opcode=SPECIAL_FGBG_2
	SPECIAL_WHITE            = 0xFD // count=1,opcode=WHITE
	SPECIAL_BLACK            = 0xFE // count=1,opcode=BLACK
)

var codeMap = map[uint8]string{
	0x00: "REGULAR_BG_RUN",
	0x01: "REGULAR_FG_RUN",
	0x02: "REGULAR_FGBG_IMAGE",
	0x03: "REGULAR_COLOR_RUN",
	0x04: "REGULAR_COLOR_IMAGE",
	0x0C: "LITE_SET_FG_FG_RUN",
	0x0D: "LITE_SET_FG_FGBG_IMAGE",
	0x0E: "LITE_DITHERED_RUN",
	0xF0: "MEGA_MEGA_BG_RUN",
	0xF1: "MEGA_MEGA_FG_RUN",
	0xF2: "MEGA_MEGA_FGBG_IMAGE",
	0xF3: "MEGA_MEGA_COLOR_RUN",
	0xF4: "MEGA_MEGA_COLOR_IMAGE",
	0xF6: "MEGA_MEGA_SET_FG_RUN",
	0xF7: "MEGA_MEGA_SET_FGBG_IMAGE",
	0xF8: "MEGA_MEGA_DITHERED_RUN",
	0xF9: "SPECIAL_FGBG_1",
	0xFA: "SPECIAL_FGBG_2",
	0xFD: "SPECIAL_WHITE",
	0xFE: "SPECIAL_BLACK",
}

var colorWhiteMap = map[int]uint32{8: 0xFF, 15: 0x7FFF, 16: 0xFFFF, 24: 0xFFFFFF}

// 获取白色像素
func getColorWhite(bpp int) uint32 {
	if white, ok := colorWhiteMap[bpp]; ok {
		return white
	} else {
		core.Throwf("invalid bpp: %v", bpp)
		return 0
	}
}

var pixelSizeMap = map[int]int{8: 1, 15: 2, 16: 2, 24: 3}

func getPixelSize(bpp int) int {
	if size, ok := pixelSizeMap[bpp]; ok {
		return size
	} else {
		core.Throwf("invalid bpp: %v", bpp)
		return 0
	}
}

// 获取黑色像素
func getColorBlack() uint32 {
	return 0x000000
}

// extract codeId
func extractCodeId(header uint8) uint8 {
	switch header & 0xF0 {
	case 0xF0: // mega&special
		return header
	case 0xC0, 0xD0, 0xE0: // lite form
		return header >> 4
	default:
		return header >> 5
	}
}

func extractRunLength(code, header uint8, r io.Reader) int {
	switch code {
	case REGULAR_FGBG_IMAGE: // FOM
		if runLength := header & g_MaskRegularRunLength; runLength != 0 {
			return int(runLength) << 3
		} else {
			return int(ReadByte(r)) + 1
		}
	case LITE_SET_FG_FGBG_IMAGE: // FOM_SET
		if runLength := header & g_MaskLiteRunLength; runLength != 0 {
			return int(runLength) << 3
		} else {
			return int(ReadByte(r)) + 1
		}
	case REGULAR_BG_RUN, REGULAR_FG_RUN, REGULAR_COLOR_RUN, REGULAR_COLOR_IMAGE: //FILL,MIX,COLOR,COPY
		if runLength := header & g_MaskRegularRunLength; runLength != 0 {
			return int(runLength)
		} else {
			return int(ReadByte(r)) + 32
		}
	case LITE_DITHERED_RUN, LITE_SET_FG_FG_RUN: // BICOLOR,MIX_SET
		if runLength := header & g_MaskLiteRunLength; runLength != 0 {
			return int(runLength)
		} else {
			return int(ReadByte(r)) + 16
		}
	case MEGA_MEGA_BG_RUN, MEGA_MEGA_COLOR_IMAGE, MEGA_MEGA_COLOR_RUN, MEGA_MEGA_DITHERED_RUN,
		MEGA_MEGA_FG_RUN, MEGA_MEGA_FGBG_IMAGE, MEGA_MEGA_SET_FG_RUN, MEGA_MEGA_SET_FGBG_IMAGE:
		return int(ReadShortLE(r))
	case SPECIAL_FGBG_1, SPECIAL_FGBG_2:
		return 8
	default:
		//glog.Debug("[DEFAULT] Return 0")
		return 0
	}
}

// 从流中读取一个像素
func readPixel(r io.Reader, bpp int) uint32 {
	switch getPixelSize(bpp) {
	case 1:
		return uint32(ReadByte(r))
	case 2:
		return uint32(ReadShortLE(r))
	case 3:
		return uint32(ReadByte(r)) | uint32(ReadShortLE(r))<<8
	}
	core.ThrowError("invalid bpp")
	return 0
}

// 写入一个像素
func writePixel(w io.Writer, pixel uint32, bpp int) {
	switch getPixelSize(bpp) {
	case 1:
		WriteByte(w, byte(pixel&0xff))
	case 2:
		WriteShortLE(w, uint16(pixel&0xffff))
	case 3:
		WriteByte(w, byte(pixel&0xff))
		WriteShortLE(w, uint16((pixel>>8)&0xffff))
	default:
		core.Throwf("invalid bpp: %v", bpp)
	}
}

// 查找上一行相同位置的像素
func peekPixel(r *bytes.Buffer, rowDelta int, bpp int) uint32 {
	if r.Len() >= rowDelta {
		pos := r.Len() - rowDelta
		return readPixel(bytes.NewReader(r.Bytes()[pos:pos+getPixelSize(bpp)]), bpp)
	}
	return getColorBlack()
}

// 解压RLE格式Bitmap
func rleDecompress(w, h, bpp int, data []byte) image.Image {
	r := bytes.NewReader(data)
	whitePixel := getColorWhite(bpp)
	blackPixel := getColorBlack()
	fgPel := whitePixel // 背景色 --> MIX

	dest := new(bytes.Buffer)

	insertFgPel := false // for FILL
	//pixels := 0

	for r.Len() > 0 {
		codeHeader := ReadByte(r)
		code := extractCodeId(codeHeader)
		//glog.Debugf("code: %s[%v]", codeMap[code], code)
		runLength := extractRunLength(code, codeHeader, r)

		// Handle Background Run Orders.
		if code == REGULAR_BG_RUN || code == MEGA_MEGA_BG_RUN { // FILL
			//pixels += runLength
			//glog.Debugf("+++ runLength: %v, pixels: %v", runLength, pixels)

			pixel := peekPixel(dest, w*getPixelSize(bpp), bpp) // 查找上一行像素
			if insertFgPel {                                   // FILL & lastcode == FILL
				writePixel(dest, pixel^fgPel, bpp)
			} else { // FILL & lastcode != FILL
				writePixel(dest, pixel, bpp)
			}
			runLength--

			for ; runLength > 0; runLength-- {
				pixel := peekPixel(dest, w*getPixelSize(bpp), bpp) // 查找上一行像素
				writePixel(dest, pixel, bpp)
			}

			//glog.Debugf("--- dest pixels: %v", dest.Len()/getPixelSize(bpp))
			insertFgPel = true // set last opcode=FILL
			continue
		}

		insertFgPel = false // set last opcode != FILL

		// change 背景色 mix
		switch code {
		case LITE_SET_FG_FG_RUN, MEGA_MEGA_SET_FG_RUN, // MIX_SET
			LITE_SET_FG_FGBG_IMAGE, MEGA_MEGA_SET_FGBG_IMAGE: // FOM_SET
			fgPel = readPixel(r, bpp)
			//glog.Debugf(" -> change fgPel: %x", fgPel)
		}

		// Process
		switch code {

		// Handle Foreground Run Orders.
		case REGULAR_FG_RUN, MEGA_MEGA_FG_RUN, LITE_SET_FG_FG_RUN, MEGA_MEGA_SET_FG_RUN: // MIX,MIX,MIX_SET,MIX_SET
			//pixels += runLength
			//glog.Debugf("+++ runLength: %v, pixels: %v", runLength, pixels)

			for ; runLength > 0; runLength-- {
				pixel := peekPixel(dest, w*getPixelSize(bpp), bpp) // 查找上一行像素
				writePixel(dest, pixel^fgPel, bpp)
			}

		// Handle Dithered Run Orders.
		case LITE_DITHERED_RUN, MEGA_MEGA_DITHERED_RUN: // BICOLOR
			//pixels += runLength * 2
			//glog.Debugf("+++ runLength: %v, pixels: %v", runLength, pixels)

			pixelA := readPixel(r, bpp)
			pixelB := readPixel(r, bpp)

			for ; runLength > 0; runLength-- {
				writePixel(dest, pixelA, bpp)
				writePixel(dest, pixelB, bpp)
			}

		// Handle Color Run Orders.
		case REGULAR_COLOR_RUN, MEGA_MEGA_COLOR_RUN: // COLOR
			//pixels += runLength
			//glog.Debugf("+++ runLength: %v, pixels: %v", runLength, pixels)

			pixelA := readPixel(r, bpp)
			for ; runLength > 0; runLength-- {
				writePixel(dest, pixelA, bpp)
			}

		// Handle Color Image Orders.
		case REGULAR_COLOR_IMAGE, MEGA_MEGA_COLOR_IMAGE: // COPY
			//pixels += runLength
			//glog.Debugf("+++ runLength: %v, pixels: %v", runLength, pixels)

			readBytes := ReadBytes(r, int(runLength)*getPixelSize(bpp))
			WriteBytes(dest, readBytes)

		// Handle Foreground/Background Image Orders.
		case REGULAR_FGBG_IMAGE, MEGA_MEGA_FGBG_IMAGE, LITE_SET_FG_FGBG_IMAGE, MEGA_MEGA_SET_FGBG_IMAGE: //FOM,FOM,FOM_SET,FOM_SET
			//pixels += runLength
			//glog.Debugf("+++ runLength: %v, pixels: %v", runLength, pixels)

			for ; runLength > 0; runLength -= 8 {
				bitmask := ReadByte(r) // => mask
				cBits := 8
				if runLength < 8 {
					cBits = runLength
				}
				for ; cBits > 0; cBits-- {
					pixel := peekPixel(dest, w*getPixelSize(bpp), bpp) // 查找上一行像素
					if bitmask&0x1 > 0 {
						pixel ^= fgPel // FIXME
					}
					writePixel(dest, pixel, bpp)
					bitmask >>= 1
				}
			}
		// Handle Special Order 1.
		case SPECIAL_FGBG_1:
			//pixels += 8
			//glog.Debugf("+++ runLength: %v, pixels: %v", "-", pixels)

			cBits := 8
			bitmask := g_MaskSpecialFgBg1
			for ; cBits > 0; cBits-- {
				pixel := peekPixel(dest, w*getPixelSize(bpp), bpp) // 查找上一行像素
				if bitmask&0x1 > 0 {
					pixel ^= fgPel // FIXME
				}
				writePixel(dest, pixel, bpp)
				bitmask >>= 1
			}

		// Handle Special Order 2.
		case SPECIAL_FGBG_2:
			//pixels += 8
			//glog.Debugf("+++ runLength: %v, pixels: %v", "-", pixels)

			cBits := 8
			bitmask := g_MaskSpecialFgBg2
			for ; cBits > 0; cBits-- {
				pixel := peekPixel(dest, w*getPixelSize(bpp), bpp) // 查找上一行像素
				if bitmask&0x1 > 0 {
					pixel ^= fgPel // FIXME
				}
				writePixel(dest, pixel, bpp)
				bitmask >>= 1
			}

		// Handle White Order.
		case SPECIAL_WHITE:
			//pixels += 1
			//glog.Debugf("+++ runLength: %v, pixels: %v", "-", pixels)

			writePixel(dest, whitePixel, bpp)

			// Handle Black Order.
		case SPECIAL_BLACK:
			//pixels += 1
			//glog.Debugf("+++ runLength: %v, pixels: %v", "-", pixels)

			writePixel(dest, blackPixel, bpp)

		default:
			core.ThrowError("invalid code")
		}
	}
	return rgb565ToImage(w, h, bpp, dest.Bytes())
}

func rgb565ToImage(w int, h int, bpp int, data []byte) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	r := bytes.NewReader(data)
	for y := 1; y <= h; y++ {
		for x := 0; x < w; x++ {
			pixel := readPixel(r, 16)
			// RGB565
			r := uint8((pixel&0xF800)>>11) << 3
			g := uint8((pixel&0x7E0)>>5) << 2
			b := uint8(pixel&0x1F) << 3
			img.Set(x, h-y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
	return img
}

// LoadRLE 加载RLE格式的Bitmap数据
func (m *BitMap) LoadRLE(option *Option) *BitMap {
	m.Image = rleDecompress(option.Width, option.Height, option.BitPerPixel, option.Data)
	return m
}
