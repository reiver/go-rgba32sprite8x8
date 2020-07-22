package rgba32sprite8x8_test

import (
	"github.com/reiver/go-rgba32sprite8x8"

	"image/draw"

	"testing"
)

func TestSlice_drawimage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum draw.Image = rgba32sprite8x8.Slice{}

	if nil == datum {
		t.Errorf("This should never happen")
	}
}
