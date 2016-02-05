package faker

import "testing"

func TestFakerName(t *testing.T) {
	f := NewDefault()
	p := f.Name()

	if len(p) == 0 {
		t.Error("Should return a full name")
	}
}

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

func TestFakerLastName(t *testing.T) {
	f := NewDefault()
	name := f.LastName()
	assertElementInSlice(t, name, f.CurrentLocale().LastName)
}
