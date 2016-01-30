package faker

// AppName returns a faker application name
func (f Faker) AppName() string {
	return randomElement(f.CurrentLocale().AppNames)
}

// AppVersion returns a fake three digit version MAJOR MINOR and PATCH
func (f Faker) AppVersion() string {
	return format("#.#.#")
}

func (f Faker) AppAuthor() string {
	return "John Doe"
}
