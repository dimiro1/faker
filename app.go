package faker

// AppName returns a faker application name
func (f Faker) AppName() string {
	return randomElement(f.CurrentLocale().AppNames)
}

// AppVersion returns a fake three digit version MAJOR MINOR and PATCH
func (f Faker) AppVersion() string {
	return format(randomElement(appVersion))
}

// AppAuthor returns a Author Name
func (f Faker) AppAuthor() string {
	return f.Name()
}
