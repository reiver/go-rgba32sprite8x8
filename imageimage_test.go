package rgba32sprite8x8_test

import (
	"github.com/reiver/go-rgba32sprite8x8"

	"image"

	"testing"
)

func TestSlice_imageimage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum image.Image = rgba32sprite8x8.Slice{}

	if nil == datum {
		t.Errorf("This should never happen")
	}
}
