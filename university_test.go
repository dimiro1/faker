package faker

import "testing"

func TestFakerUniversityName(t *testing.T) {
	f := NewDefault()
	u := f.UniversityName()
	assertElementInSlice(t, u, f.CurrentLocale().Universities)
}
