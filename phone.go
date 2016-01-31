package faker

// PhoneNumber returns a phone number
func (f Faker) PhoneNumber() string {
	return format(randomElement(f.CurrentLocale().PhoneNumbers))
}

// CellPhoneNumber returns a cellphone number
func (f Faker) CellPhoneNumber() string {
	return format(randomElement(f.CurrentLocale().PhoneNumbers))
}

// PhoneAreaCode returns an area code
func (f Faker) PhoneAreaCode() string {
	return randomElement(f.CurrentLocale().AreaCodes)
}

// PhoneExchangeCode returns an exchange code
func (f Faker) PhoneExchangeCode() string {
	return randomElement(f.CurrentLocale().ExchangeCodes)
}
