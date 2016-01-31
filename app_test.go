package faker

import "testing"

func TestFakerAppName(t *testing.T) {
	f := NewDefault()
	u := f.AppName()
	assertElementInSlice(t, u, f.CurrentLocale().AppNames)
}

func TestFakerAppVersion(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^\\d\\.\\d(\\.\\d)?(-(alpha|beta|rc1|rc2|release))?$", f.AppVersion())
}
