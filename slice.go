package rgba32sprite8x8

import (
	"github.com/reiver/go-rgba32"

	"image"
	"image/color"
)

type Slice []uint8

func (receiver Slice) At(x, y int) color.Color {
	if nil == receiver {
		return color.RGBA{}
	}

	if ! (image.Point{X:y,Y:y}).In(receiver.Bounds()) {
		return color.RGBA{}
	}

	low  := receiver.PixOffset(x,y)
	high := low + Depth

	p := receiver[low:high]

	return rgba32.Slice(p)
}

func (receiver Slice) Bounds() image.Rectangle {
	const x = 0
	const y = 0

	return image.Rectangle{
		Min:image.Point{
			X: x,
			Y: y,
		},
		Max:image.Point{
			X: x+Width,
			Y: y+Height,
		},
	}
}

func (receiver Slice) ColorModel() color.Model {
	return color.RGBAModel
}

func (receiver Slice) PixOffset(x, y int) int {
	rect := receiver.Bounds()

	xx := x - rect.Min.X
	yy := y - rect.Min.Y

	return yy*Width*Depth + xx*Depth
}

func (receiver Slice) Set(x, y int, c color.Color) {
	if nil == receiver {
		return
	}

	cc := receiver.At(x,y)
	if nil == cc {
		return
	}

	rgba, casted := cc.(rgba32.Slice)
	if !casted {
		return
	}
	if rgba32.ByteSize != len(rgba) {
		return
	}

	rr,gg,bb,aa := c.RGBA()

	r := uint8((rr*0xff)/0xffff)
	g := uint8((gg*0xff)/0xffff)
	b := uint8((bb*0xff)/0xffff)
	a := uint8((aa*0xff)/0xffff)

	rgba[rgba32.OffsetRed]   = r
	rgba[rgba32.OffsetGreen] = g
	rgba[rgba32.OffsetBlue]  = b
	rgba[rgba32.OffsetAlpha] = a
}
