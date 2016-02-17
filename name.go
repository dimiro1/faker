package faker

// Name returns a Male or Female full name
func (f Faker) Name() string {
	return template("Name", randomElement(f.CurrentLocale().NamesFormats), f)
}

// FirstName returns a Male or Female first name
func (f Faker) FirstName() string {
	return randomElement(f.CurrentLocale().FirstNames)
}

// LastName returns a Male or Female last name
func (f Faker) LastName() string {
	return randomElement(f.CurrentLocale().LastName)
}

// Gender returns a gender name
func (f Faker) Gender() string {
	return randomElement(f.CurrentLocale().Gender)
}
