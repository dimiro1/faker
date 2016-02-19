package faker

import (
	"fmt"
	"strings"

	"github.com/dimiro1/faker/defaults"
)

type ImageOptions struct {
	Size   string `default:"300x300"`
	Format string `default:".jpg"`
}

func (f Faker) AvatarImage() string {
	return "AvatarImage"
}

func (f Faker) AvatarImageWithOptions(o ImageOptions) string {
	return "AvatarImageWithOptions"
}

// PlaceholderImage returns a Placeholder image
func (f Faker) PlaceholderImage() string {
	width := randomIntBetween(100, 400)
	height := randomIntBetween(100, 400)

	return fmt.Sprintf("http://placehold.it/%dx%d", width, height)
}

// PlaceholderImageWithOptions returns a Placeholder image
func (f Faker) PlaceholderImageWithOptions(o ImageOptions) string {
	defaults.Apply(&o)

	formats := []string{".jpeg", ".jpg", ".gif", ".png"}
	valid := false

	for _, format := range formats {
		if o.Format == format {
			valid = true
		}
	}

	if !valid {
		panic(fmt.Sprintf("images: %s is an invalid format, valid formats are %s", o.Format, formats))
	}

	size := strings.Split(o.Size, "x")

	if len(size) != 2 {
		panic(fmt.Sprintf("images: %s is an invalid size, valid format is 300x300", o.Size))
	}

	width := size[0]
	height := size[1]

	return fmt.Sprintf("http://placehold.it/%sx%s%s", width, height, o.Format)
}
