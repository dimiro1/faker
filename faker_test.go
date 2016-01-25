package faker

import "testing"

func TestNewFakerWithInvalidLocale(t *testing.T) {
	_, err := NewFaker(Options{Locale: "invalid"})

	if err == nil {
		t.Error("Should panic with invalid Locale")
	}
}
