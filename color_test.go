package faker

import "testing"

func TestFakerColorName(t *testing.T) {
	f := NewDefault()
	l := f.ColorName()
	assertElementInSlice(t, l, colorNames)
}

func TestFakerHexColor(t *testing.T) {
	f := NewDefault()
	c := f.HexColor()
	assertStringRegexp(t, "[0-9a-f]{8}", c)
}
