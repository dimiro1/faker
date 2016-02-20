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

type AvatarOptions struct {
	Size   string `default:"300"`
	Format string `default:".jpg"`
	Email  string
}

// AvatarImage returns a Avatar image
func (f Faker) AvatarImage() string {
	width := randomIntBetween(100, 400)

	return f.AvatarImageWithOptions(AvatarOptions{
		Size:   fmt.Sprintf("%d", width),
		Format: ".jpg",
		Email:  f.Email(),
	})
}

// AvatarImageWithOptions returns a AvatarImage image
func (f Faker) AvatarImageWithOptions(o AvatarOptions) string {
	defaults.Apply(&o)

	if o.Email == "" {
		panic(fmt.Sprintf("images: %s is empty", o.Email))
	}

	formats := []string{".jpg", ".png"}
	valid := false

	for _, format := range formats {
		if o.Format == format {
			valid = true
		}
	}

	if !valid {
		panic(fmt.Sprintf("images: %s is an invalid format, valid formats are %s", o.Format, formats))
	}

	return fmt.Sprintf("http://api.adorable.io/avatars/%s/%s%s", o.Size, o.Email, o.Format)
}

// PlaceholderImage returns a Placeholder image
func (f Faker) PlaceholderImage() string {
	width := randomIntBetween(100, 400)
	height := randomIntBetween(100, 400)

	return f.PlaceholderImageWithOptions(ImageOptions{
		Size: fmt.Sprintf("%dx%d", width, height),
	})
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
