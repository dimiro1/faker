package faker

// Name returns a Male or Female full name
func (f Faker) Name() string {
	return template("Name", randomElement(f.CurrentLocale().Names), f)
}

// FirstName returns a Male or Female first name
func (f Faker) FirstName() string {
	return randomElement(f.CurrentLocale().FirstNames)
}

func (f Faker) LastName() string {
	return "Neto"
}

func (f Faker) Prefix() string {
	return "Sr."
}

func (f Faker) Suffix() string {
	return "III"
}

func (f Faker) Title() string {
	return "Legacy Creative Director"
}

// Gender returns a gender name
func (f Faker) Gender() string {
	return randomElement(f.CurrentLocale().Gender)
}
