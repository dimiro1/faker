package faker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dimiro1/faker/locales"
)

// DefaultLocale is a constant that defines the default locale used in the default configuration.
const DefaultLocale = "en"

// Faker is the main API
// it also holds some config
type Faker struct {
	CurrentLocaleString string
	Locales             map[string]locales.Locale
}

// Options is used to pass configuration options to Faker
type Options struct {
	Locale string
	Seed   int64
}

// CurrentLocale returns the current locale object
func (f Faker) CurrentLocale() locales.Locale {
	return f.Locales[f.CurrentLocaleString]
}

// NewDefault creates a new Faker with default configuration
func NewDefault() Faker {
	f, err := NewForLocale(DefaultLocale)

	if err != nil {
		panic(fmt.Sprintf("NewDefaultFaker: NewFaker returned an error with valid options %+v\n", err))
	}

	return f
}

// NewForLocale creates a new Faker with the specified locale and random seed
func NewForLocale(locale string) (Faker, error) {
	return New(Options{Locale: locale, Seed: time.Now().UTC().UnixNano()})
}

// New creates a new Faker with configuration passed in options
func New(options Options) (Faker, error) {

	// Validate Locale
	if !locales.IsValid(options.Locale) {
		return Faker{}, fmt.Errorf("NewFaker: %s is an invalid locale", options.Locale)
	}

	// Seeding the random
	rand.Seed(options.Seed)

	return Faker{
		CurrentLocaleString: options.Locale,
		Locales:             locales.Supported(),
	}, nil
}
