package faker

import "testing"

func TestFakerAppName(t *testing.T) {
	f := NewDefault()
	u := f.AppName()
	assertElementInSlice(t, u, f.CurrentLocale().AppNames)
}

func TestFakerAppVersion(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^\\d\\.\\d\\.\\d$", f.AppVersion())
}
