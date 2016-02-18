package faker

import "testing"

func TestFakerPlaceholderImage(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^http://placehold.it/\\d+x\\d+$", f.PlaceholderImage())
}
