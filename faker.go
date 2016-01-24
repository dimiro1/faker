package faker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dimiro1/faker/locales"
)

type Faker struct {
	Locale  string
	Locales map[string]locales.Locale
}

type FakerOptions struct {
	Locale string
	Seed   int64
}

func NewDefaultFaker() Faker {
	f, err := NewFaker(FakerOptions{
		Locale: "en",
		Seed:   time.Now().UTC().UnixNano(),
	})

	if err != nil {
		panic(fmt.Sprintf("This should never happen: %+v\n", err))
	}

	return f
}

func NewFaker(options FakerOptions) (Faker, error) {
	// Validate Locale
	// Validate Seed

	// Seeding the random
	rand.Seed(options.Seed)

	return Faker{
		Locale:  options.Locale,
		Locales: locales.Supported(),
	}, nil
}
