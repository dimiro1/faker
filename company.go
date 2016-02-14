package faker

// CompanyName returns a company name
func (f Faker) CompanyName() string {
	return template("CompanyName", randomElement(f.CurrentLocale().CompanyNamesFormats), f)
}

// CompanySuffix returns a company suffix such as 'corp', 'LLC'
func (f Faker) CompanySuffix() string {
	return randomElement(f.CurrentLocale().CompanySuffixes)
}

func (f Faker) CompanyCatchPhrase() string {
	return "CompanyCatchPhrase"
}

func (f Faker) CompanyBuzzword() string {
	return "CompanyBuzzword"
}

func (f Faker) CompanyEIN() string {
	return "CompanyEIN"
}

// CompanyLogo returns a fake URL company name
func (f Faker) CompanyLogo() string {
	return randomElement(companyLogos)
}
