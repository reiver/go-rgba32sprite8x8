package rgba32sprite8x8_test

import (
	"github.com/reiver/go-rgba32sprite8x8"

	"github.com/reiver/go-rgba32"

	"image"
	"math/rand"
	"time"

	"testing"
)

func TestSlice_At_simple(t *testing.T) {

	for testNumber:=0; testNumber<10; testNumber++ {

		var buffer [rgba32sprite8x8.ByteSize]uint8
		var sprite rgba32sprite8x8.Slice = rgba32sprite8x8.Slice(buffer[:])

		for offset:=0; offset<len(buffer); offset++ {
			buffer[offset] = uint8(offset)
		}

		{
			var img image.Image = sprite

			for y:=0;y<rgba32sprite8x8.Height;y++ {
				for x:=0;x<rgba32sprite8x8.Width;x++ {
					c := img.At(x,y)
					rgba := c.(rgba32.Slice)

					aR := rgba[0]
					aG := rgba[1]
					aB := rgba[2]
					aA := rgba[3]

					i := y*8*4 + x*4

					eR := uint8(i  )
					eG := uint8(i+1)
					eB := uint8(i+2)
					eA := uint8(i+3)

					if eR != aR || eG != aG || eB != aB || eA != aA {
						t.Errorf("For test #%d, the actual color is not what was expected.", testNumber)
						t.Logf("EXPECTED (r,g,b,a) = (%d,%d,%d,%d)", eR, eG, eB, eA)
						t.Logf("ACTUAL   (r,g,b,a) = (%d,%d,%d,%d)", aR, aG, aB, aA)
					}
				}
			}
		}
	}
}

func TestSlice_At_randomized(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for testNumber:=0; testNumber<10; testNumber++ {

		var buffer [rgba32sprite8x8.ByteSize]uint8
		var sprite rgba32sprite8x8.Slice = rgba32sprite8x8.Slice(buffer[:])

		for offset:=0; offset<len(buffer); offset++ {
			buffer[offset] = uint8(randomness.Intn(256))
		}

		{
			var img image.Image = sprite

			for y:=0;y<rgba32sprite8x8.Height;y++ {
				for x:=0;x<rgba32sprite8x8.Width;x++ {
					c := img.At(x,y)
					rgba := c.(rgba32.Slice)

					aR := rgba[0]
					aG := rgba[1]
					aB := rgba[2]
					aA := rgba[3]

					i := y*8*4 + x*4

					eR := buffer[i  ]
					eG := buffer[i+1]
					eB := buffer[i+2]
					eA := buffer[i+3]

					if eR != aR || eG != aG || eB != aB || eA != aA {
						t.Errorf("For test #%d, the actual color is not what was expected.", testNumber)
						t.Logf("EXPECTED (r,g,b,a) = (%d,%d,%d,%d)", eR, eG, eB, eA)
						t.Logf("ACTUAL   (r,g,b,a) = (%d,%d,%d,%d)", aR, aG, aB, aA)
					}
				}
			}
		}
	}
}
