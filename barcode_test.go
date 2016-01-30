package faker

import "testing"

func TestFakerEAN13(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^\\d{13}$", f.EAN13())
}
