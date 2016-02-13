package locales

// Locale is a struct that hold the common names from each lang
type Locale struct {
	Code string

	AddressesFormats          []string
	AppNames                  []string
	AreaCodes                 []string
	BookGenres                []string
	BookPublishersFormats     []string
	BookPublishersSuffixes    []string
	BuildingNumbers           []string
	CityNamesFormats          []string
	CityPrefix                []string
	CitySuffix                []string
	CompanyNamesFormats       []string
	CompanySuffixes           []string
	CountryNames              []string
	ExchangeCodes             []string
	FirstNames                []string
	Gender                    []string
	LastName                  []string
	NamesFormats              []string
	PhoneNumbersFormats       []string
	SecondaryAddressesFormats []string
	StateAbbr                 []string
	StateNames                []string
	StreetAddressesFormats    []string
	StreetNamesFormats        []string
	StreetSuffixes            []string
	Universities              []string
	ZipFormats                []string
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
