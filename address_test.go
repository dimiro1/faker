package faker

import "testing"

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

func TestFakerCityPrefix(t *testing.T) {
	f := NewDefault()
	c := f.CityPrefix()
	assertElementInSlice(t, c, f.CurrentLocale().CityPrefix)
}

func TestFakerCitySuffix(t *testing.T) {
	f := NewDefault()
	c := f.CitySuffix()
	assertElementInSlice(t, c, f.CurrentLocale().CitySuffix)
}

func TestFakerLatitude(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^(\\+|-)?(?:90(?:(?:\\.0{1,6})?)|(?:[0-9]|[1-8][0-9])(?:(?:\\.[0-9]{1,6})?))$", f.Latitude())
}

func TestFakerLongitude(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^(\\+|-)?(?:180(?:(?:\\.0{1,6})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\\.[0-9]{1,6})?))$", f.Longitude())
}

func TestFakerZipCode(t *testing.T) {
	f := NewDefault()
	p := f.ZipCode()

	if len(p) == 0 {
		t.Error("Should return a zip code")
	}
}
