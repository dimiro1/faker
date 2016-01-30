package faker

// UniversityName returns a name of a University
func (f Faker) UniversityName() string {
	return randomElement(f.CurrentLocale().Universities)
}
