package locales

// Locale is a struct that hold the common names from each lang
type Locale struct {
	Code string

	AppNames      []string
	AreaCodes     []string
	CityPrefix    []string
	CitySuffix    []string
	CityNames     []string
	CountryNames  []string
	ExchangeCodes []string
	StateAbbr     []string
	StateNames    []string
	Universities  []string
	PhoneNumbers  []string
}

// IsValid returns true if the given locale is a valid locale
func IsValid(locale string) bool {
	_, ok := Supported()[locale]
	return ok
}

// Supported returns a map of supported Locales
func Supported() map[string]Locale {
	return map[string]Locale{
		"en":    en_locale(),
		"pt-br": pt_br_locale(),
	}
}
