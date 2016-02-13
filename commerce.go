package faker

// StoreDepartment returns a fake department name
func (f Faker) StoreDepartment() string {
	return randomElement(f.CurrentLocale().StoreDepartments)
}

// ProductSimpleName returns a fake product name
func (f Faker) ProductSimpleName() string {
	return randomElement(f.CurrentLocale().ProductNames)
}

// ProductName returns a fake product name
func (f Faker) ProductName() string {
	return template("ProductName", randomElement(f.CurrentLocale().ProductNamesFormats), f)
}

// ProductAdjective returns a fake product adjective
func (f Faker) ProductAdjective() string {
	return randomElement(f.CurrentLocale().ProductAdjectives)
}

// ProductPrice returns a fake product price between 0 and 5000
func (f Faker) ProductPrice() float64 {
	return randomFloatBetween(0, 5000)
}
