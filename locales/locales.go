package locales

type Locale struct {
	CountryNames []string
	CityPrefix   []string
	CitySuffix   []string
}

func Supported() map[string]Locale {
	return map[string]Locale{
		"en":    newENLocale(),
		"pt-br": newPTBRLocale(),
	}
}
