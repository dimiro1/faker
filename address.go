package faker

import "fmt"

// City returns a city name
func (f Faker) City() string {
	return randomElement(f.CurrentLocale().CityNames)
}

// StreetName return a street name
func (f Faker) StreetName() string {
	return template("StreetName", randomElement(f.CurrentLocale().StreetNames), f)
}

// StreetAddress return a full street address
func (f Faker) StreetAddress() string {
	return template("StreetAddress", randomElement(f.CurrentLocale().StreetAddress), f)
}

func (f Faker) SecondaryAddress() string {
	return "SecondaryAddress"
}

// BuildingNumber return a number of a building
func (f Faker) BuildingNumber() string {
	return format(randomElement(f.CurrentLocale().BuildingNumbers))
}

// ZipCode must return a zip code
func (f Faker) ZipCode() string {
	return randomElement(f.CurrentLocale().Zip)
}

// StreetSuffix return a Street Suffix
func (f Faker) StreetSuffix() string {
	return randomElement(f.CurrentLocale().StreetSuffixes)
}

// CityPrefix return a City Prefix
func (f Faker) CityPrefix() string {
	return randomElement(f.CurrentLocale().CityPrefix)
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
	return randomElement(countryCodes)
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
