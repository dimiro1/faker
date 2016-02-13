package faker

import "testing"

func TestFakerBookPublisherSuffix(t *testing.T) {
	f := NewDefault()
	u := f.BookPublisherSuffix()
	assertElementInSlice(t, u, f.CurrentLocale().BookPublishersSuffixes)
}
