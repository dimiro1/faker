package faker

import "testing"

func TestFakerCountry(t *testing.T) {
	f := NewDefault()
	country := f.Country()
	assertElementInSlice(t, country, f.CurrentLocale().CountryNames)
}

func TestFakerState(t *testing.T) {
	f := NewDefault()
	state := f.State()
	assertElementInSlice(t, state, f.CurrentLocale().StateNames)
}

func TestFakerStateAbbr(t *testing.T) {
	f := NewDefault()
	abbr := f.StateAbbr()
	assertElementInSlice(t, abbr, f.CurrentLocale().StateAbbr)
}
