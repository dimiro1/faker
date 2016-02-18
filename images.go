package faker

import "fmt"

type ImageOptions struct {
	Size   string `default:"300x300"`
	Format string `default:"jpg"`
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

func (f Faker) PlaceholderImageWithOptions(o ImageOptions) string {
	return "AvatarImageWithOptions"
}
