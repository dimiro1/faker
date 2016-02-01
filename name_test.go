package faker

import "testing"

func TestFakerFirstName(t *testing.T) {
	f := NewDefault()
	name := f.FirstName()
	assertElementInSlice(t, name, f.CurrentLocale().FirstNames)
}

func TestFakerGender(t *testing.T) {
	f := NewDefault()
	g := f.Gender()
	assertElementInSlice(t, g, f.CurrentLocale().Gender)
}
