package faker

import "testing"

func TestNewForLocale(t *testing.T) {
	f, err := NewForLocale(DefaultLocale)

	if err != nil {
		t.Error(err)
	}

	if f.CurrentLocale().Code != DefaultLocale {
		t.Errorf("The default locale should be %s", DefaultLocale)
	}
}

func TestNewFakerWithInvalidLocale(t *testing.T) {
	_, err := New(Options{Locale: "invalid"})

	if err == nil {
		t.Error("Should panic with invalid Locale")
	}
}

func TestFakerCurrentLocale(t *testing.T) {
	f := NewDefault()

	if f.CurrentLocale().Code != DefaultLocale {
		t.Errorf("The default locale should be %s", DefaultLocale)
	}
}
