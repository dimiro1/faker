package faker

// StoreDepartment returns a fake department name
func (f Faker) StoreDepartment() string {
	return randomElement(f.CurrentLocale().StoreDepartments)
}

func (f Faker) ProductName() string {
	return "ProductName"
}

// ProductPrice returns a fake product price between 0 and 5000
func (f Faker) ProductPrice() float64 {
	return randomFloatBetween(0, 5000)
}
