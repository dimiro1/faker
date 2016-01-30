package faker

import "testing"

func TestFakerColorName(t *testing.T) {
	f := NewDefault()
	l := f.ColorName()
	assertElementInSlice(t, l, colorNames)
}
