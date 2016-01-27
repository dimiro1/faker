package faker

import "testing"

func TestFakerCountry(t *testing.T) {
	f := NewDefault()
	country := f.Country()
	assertElementInSlice(t, country, f.CurrentLocale().CountryNames)
}
