package faker

import "testing"

func TestFakerCity(t *testing.T) {
	f := NewDefault()
	city := f.City()
	assertElementInSlice(t, city, f.CurrentLocale().CityNames)
}

func TestFakerCountry(t *testing.T) {
	f := NewDefault()
	country := f.Country()
	assertElementInSlice(t, country, f.CurrentLocale().CountryNames)
}

func TestFakerCountryCode(t *testing.T) {
	f := NewDefault()
	code := f.CountryCode()
	assertElementInSlice(t, code, countryCodes)
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

func TestFakerLatitude(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^(\\+|-)?(?:90(?:(?:\\.0{1,6})?)|(?:[0-9]|[1-8][0-9])(?:(?:\\.[0-9]{1,6})?))$", f.Latitude())
}

func TestFakerLongitude(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^(\\+|-)?(?:180(?:(?:\\.0{1,6})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\\.[0-9]{1,6})?))$", f.Longitude())
}
