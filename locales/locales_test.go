package locales

import "testing"

func TestIsValid(t *testing.T) {
	if IsValid("invalid") {
		t.Error("'invalid' is not a valid locale")
	}

	if !IsValid("en") {
		t.Error("'en' is a valid locale")
	}
}
