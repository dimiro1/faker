package faker

import "fmt"

func (f Faker) City() string {
	return "City"
}

func (f Faker) StreetName() string {
	return "StreetName"
}

func (f Faker) StreetAddress() string {
	return "StreetAddress"
}

func (f Faker) SecondaryAddress() string {
	return "SecondaryAddress"
}

func (f Faker) BuildingNumber() string {
	return "BuildingNumber"
}

func (f Faker) ZipCode() string {
	return "ZipCode"
}

func (f Faker) Zip() string {
	return "Zip"
}

func (f Faker) StreetSuffix() string {
	return "StreetSuffix"
}

func (f Faker) CityPrefix() string {
	return "CityPrefix"
}

// State returns a state name
func (f Faker) State() string {
	return randomElement(f.CurrentLocale().StateNames)
}

// StateAbbr returns a state abbrewviation
func (f Faker) StateAbbr() string {
	return randomElement(f.CurrentLocale().StateAbbr)
}

// Country returns a country name
func (f Faker) Country() string {
	return randomElement(f.CurrentLocale().CountryNames)
}

// CountryCode returns a ISO Country code
func (f Faker) CountryCode() string {
	return randomElement(f.CurrentLocale().CountryCodes)
}

// Latitude returns an earth latitude
func (f Faker) Latitude() string {
	return fmt.Sprintf("%f", geoCordinate()/2)
}

// Longitude returns an earth longitude
func (f Faker) Longitude() string {
	return fmt.Sprintf("%f", geoCordinate())
}

func geoCordinate() float64 {
	return randomFloatBetween(-180000000, 180000000) / 1000000.0
}
