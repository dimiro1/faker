package faker

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

func (f Faker) State() string {
	return "State"
}

func (f Faker) StateAbbr() string {
	return "StateAbbr"
}

// Country returns a country name
func (f Faker) Country() string {
	return anyFromSlice(f.CurrentLocale().CountryNames)
}

func (f Faker) CountryCode() string {
	return "CountryCode"
}

func (f Faker) Latitude() string {
	return "Latitude"
}

func (f Faker) Longitude() string {
	return "Longitude"
}
