package faker

import "fmt"

// Address returns a full address
func (f Faker) Address() string {
	return template("Address", randomElement(f.CurrentLocale().AddressesFormats), f)
}

// City returns a city name
func (f Faker) City() string {
	return template("City", randomElement(f.CurrentLocale().CityNamesFormats), f)
}

// StreetName return a street name
func (f Faker) StreetName() string {
	return template("StreetName", randomElement(f.CurrentLocale().StreetNamesFormats), f)
}

// StreetAddress return a full street address
func (f Faker) StreetAddress() string {
	return template("StreetAddress", randomElement(f.CurrentLocale().StreetAddressesFormats), f)
}

// SecondaryAddress returns a Secondary Address
func (f Faker) SecondaryAddress() string {
	return format(template("SecondaryAddress", randomElement(f.CurrentLocale().SecondaryAddressesFormats), f))
}

// BuildingNumber return a number of a building
func (f Faker) BuildingNumber() string {
	return format(randomElement(f.CurrentLocale().BuildingNumbers))
}

// ZipCode must return a zip code
func (f Faker) ZipCode() string {
	return format(randomElement(f.CurrentLocale().ZipFormats))
}

// StreetSuffix return a Street Suffix
func (f Faker) StreetSuffix() string {
	return randomElement(f.CurrentLocale().StreetSuffixes)
}

// CityPrefix return a City Prefix
func (f Faker) CityPrefix() string {
	return randomElement(f.CurrentLocale().CityPrefix)
}

// CitySuffix return a City Suffix
func (f Faker) CitySuffix() string {
	return randomElement(f.CurrentLocale().CitySuffix)
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
