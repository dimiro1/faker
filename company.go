package faker

// CompanyName returns a company name
func (f Faker) CompanyName() string {
	return template("CompanyName", randomElement(f.CurrentLocale().CompanyNames), f)
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

func (f Faker) CompanyLogo() string {
	return "https://pigment.github.com/fake-logos/logos/medium/color/5.png"
}
