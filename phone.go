package faker

func (f Faker) PhoneNumber() string {
	return "(11) 98161-0983"
}

func (f Faker) CellPhoneNumber() string {
	return "(11) 98161-0983"
}

// PhoneAreaCode returns an area code
func (f Faker) PhoneAreaCode() string {
	return randomElement(f.CurrentLocale().AreaCodes)
}

// PhoneExchangeCode returns an exchange code
func (f Faker) PhoneExchangeCode() string {
	return randomElement(f.CurrentLocale().ExchangeCodes)
}
