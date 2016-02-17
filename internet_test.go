package faker

import "testing"

func TestUserName(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^\\w+$", f.UserName())
}
