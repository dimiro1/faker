package faker

import "testing"

func TestFakerAvatarImage(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^http://api.adorable.io/avatars/(\\d+)/(.+).jpg$", f.AvatarImage())
}

func TestFakerAvatarImageWithOptions_with_invalid_format(t *testing.T) {
	f := NewDefault()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Should panic with wrong Format argument")
		}
	}()

	f.AvatarImageWithOptions(AvatarOptions{Format: ".asas", Email: "hello@example.com"})
}

func TestFakerAvatarImageWithOptions_with_invalid_size(t *testing.T) {
	f := NewDefault()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Should panic with wrong Size argument")
		}
	}()

	f.AvatarImageWithOptions(AvatarOptions{Size: "asasacdcd"})
}

func TestFakerAvatarImageWithOptions_with_empty_email(t *testing.T) {
	f := NewDefault()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Should panic with empty email")
		}
	}()

	f.AvatarImageWithOptions(AvatarOptions{Email: ""})
}

func TestFakerPlaceholderImage(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^http://placehold.it/\\d+x\\d+\\.(jpeg|jpg|gif|png)$", f.PlaceholderImage())
}

func TestFakerPlaceholderImageWithOptions_with_defaults(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^http://placehold.it/\\d+x\\d+\\.(jpeg|jpg|gif|png)$", f.PlaceholderImageWithOptions(ImageOptions{}))
}

func TestFakerPlaceholderImageWithOptions(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^http://placehold.it/400x500.png$", f.PlaceholderImageWithOptions(ImageOptions{Size: "400x500", Format: ".png"}))
}

func TestFakerPlaceholderImageWithOptions_with_invalid_format(t *testing.T) {
	f := NewDefault()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Should panic with wrong Format argument")
		}
	}()

	f.PlaceholderImageWithOptions(ImageOptions{Format: ".asas"})
}

func TestFakerPlaceholderImageWithOptions_with_invalid_size(t *testing.T) {
	f := NewDefault()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Should panic with wrong Size argument")
		}
	}()

	f.PlaceholderImageWithOptions(ImageOptions{Size: "asasacdcd"})
}
