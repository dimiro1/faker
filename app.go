package faker

func (f Faker) AppName() string {
	return "Software X"
}

// AppVersion returns a fake three digit version MAJOR MINOR and PATCH
func (f Faker) AppVersion() string {
	return format("#.#.#")
}

func (f Faker) AppAuthor() string {
	return "John Doe"
}
