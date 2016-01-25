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
	CurrentLocale string
	Locales       map[string]locales.Locale
}

// Options is used to pass configuration options to Faker
type Options struct {
	Locale string
	Seed   int64
}

// NewDefaultFaker creates a new Faker with default configuration
func NewDefaultFaker() Faker {
	f, err := NewFaker(Options{
		Locale: DefaultLocale,
		Seed:   time.Now().UTC().UnixNano(),
	})

	if err != nil {
		panic(fmt.Sprintf("This should never happen: %+v\n", err))
	}

	return f
}

// NewFaker creates a new Faker with configuration passed in options
func NewFaker(options Options) (Faker, error) {

	// Validate Locale
	if !locales.IsValid(options.Locale) {
		return Faker{}, fmt.Errorf("%s is an invalid locale", options.Locale)
	}

	// Seeding the random
	rand.Seed(options.Seed)

	return Faker{
		CurrentLocale: options.Locale,
		Locales:       locales.Supported(),
	}, nil
}
