package rgba32sprite8x8_test

import (
	"github.com/reiver/go-rgba32sprite8x8"

	"image"

	"testing"
)

func TestSlice_Bounds(t *testing.T) {

	var img image.Image = rgba32sprite8x8.Slice(nil)

	bounds := img.Bounds()

	if expected, actual := 0, bounds.Min.X; expected != actual {
		t.Errorf("The actual min-x bound is not what was expected.")
		t.Logf("EXPECTED: %d", expected)
		t.Logf("ACTUAL:   %d", actual)
	}
	if expected, actual := 0, bounds.Min.Y; expected != actual {
		t.Errorf("The actual min-y bound is not what was expected.")
		t.Logf("EXPECTED: %d", expected)
		t.Logf("ACTUAL:   %d", actual)
	}

	if expected, actual := rgba32sprite8x8.Width, bounds.Max.X; expected != actual {
		t.Errorf("The actual max-x bound is not what was expected.")
		t.Logf("EXPECTED: %d", expected)
		t.Logf("ACTUAL:   %d", actual)
	}
	if expected, actual := rgba32sprite8x8.Height, bounds.Max.Y; expected != actual {
		t.Errorf("The actual max-y bound is not what was expected.")
		t.Logf("EXPECTED: %d", expected)
		t.Logf("ACTUAL:   %d", actual)
	}

}
