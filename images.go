package faker

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

func (f Faker) PlaceholderImage() string {
	return "Placeholder"
}

func (f Faker) PlaceholderImageWithOptions(o ImageOptions) string {
	return "AvatarImageWithOptions"
}
