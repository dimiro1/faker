package faker

import "testing"

func TestFakerBookPublisherSuffix(t *testing.T) {
	f := NewDefault()
	u := f.BookPublisherSuffix()
	assertElementInSlice(t, u, f.CurrentLocale().BookPublishersSuffixes)
}

func TestFakerBookGenre(t *testing.T) {
	f := NewDefault()
	u := f.BookGenre()
	assertElementInSlice(t, u, f.CurrentLocale().BookGenres)
}
