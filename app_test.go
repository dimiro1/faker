package faker

import "testing"

func TestFakerAppName(t *testing.T) {
	f := NewDefault()
	u := f.AppName()
	assertElementInSlice(t, u, f.CurrentLocale().AppNames)
}

func TestFakerAppAuthor(t *testing.T) {
	f := NewDefault()
	u := f.AppAuthor()

	if len(u) == 0 {
		t.Error("Should return a author name")
	}
}

func TestFakerAppVersion(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^\\d\\.\\d(\\.\\d)?(-(alpha|beta|rc1|rc2|release))?$", f.AppVersion())
}
