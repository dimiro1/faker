package faker

import "testing"

func TestFakerPlaceholderImage(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^http://placehold.it/\\d+x\\d+$", f.PlaceholderImage())
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
