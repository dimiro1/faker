package faker

import "testing"

func TestFakerAppVersion(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^\\d\\.\\d\\.\\d$", f.AppVersion())
}
