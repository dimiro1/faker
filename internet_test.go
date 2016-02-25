package faker

import "testing"

func TestFakerPassword(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^[a-f0-9]{8}$", f.Password())
}

func TestFakerPasswordWithOptions(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^[a-f0-9]{10}$", f.PasswordWithOptions(PassWordOptions{Length: 10, Alpha: true}))
}

func TestFakerPasswordWithOptions_Alpha_false(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^[0-9]{10}$", f.PasswordWithOptions(PassWordOptions{Length: 10, Alpha: false}))
}

func TestFakerUserName(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^\\w+$", f.UserName())
}

func TestFakerUserNameWithOptions(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "fulano", f.UserNameWithOptions(UserNameOptions{Name: "Fulano"}))
	assertStringRegexp(t, "fulano-de-tal", f.UserNameWithOptions(UserNameOptions{Name: "Fulano de Tal"}))
}
