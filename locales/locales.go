package locales

// Locale is a struct that hold the common names from each lang
type Locale struct {
	CountryNames []string
	CityPrefix   []string
	CitySuffix   []string
}

// IsValid returns true if the given locale is a valid locale
func IsValid(locale string) bool {
	for k := range Supported() {
		if k == locale {
			return true
		}
	}

	return false
}

// Supported returns a map of supported Locales
func Supported() map[string]Locale {
	return map[string]Locale{
		"en":    en_locale(),
		"pt-br": pt_br_locale(),
	}
}
