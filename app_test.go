package faker

import "testing"

func TestFakerAppVersion(t *testing.T) {
	f := NewDefaultFaker()
	assertStringRegexp(t, "^\\d\\.\\d\\.\\d$", f.AppVersion())
}
