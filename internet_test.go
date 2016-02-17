package faker

import "testing"

func TestFakerUserName(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^\\w+$", f.UserName())
}

func TestFakerUserNameWithOptions(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "fulano", f.UserNameWithOptions(UserNameOptions{Name: "Fulano"}))
	assertStringRegexp(t, "fulano-de-tal", f.UserNameWithOptions(UserNameOptions{Name: "Fulano de Tal"}))
}
