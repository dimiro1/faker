package faker

import (
	"strings"
	"testing"
)

func TestFakerDomainSuffix(t *testing.T) {
	f := NewDefault()
	s := f.DomainSuffix()
	assertElementInSlice(t, s, domainSuffixes)
}

func TestFakerSafeEmail(t *testing.T) {
	f := NewDefault()
	email := f.FreeEmail()
	splited := strings.Split(email, "@")

	assertStringRegexp(t, "^\\w+$", splited[0])
	assertElementInSlice(t, splited[1], freeEmailDomains)
}

func TestFakerSafeEmailWithOptions(t *testing.T) {
	f := NewDefault()
	email := f.FreeEmailWithOptions(EmailOptions{Name: "Fulano"})
	splited := strings.Split(email, "@")

	assertStringRegexp(t, "fulano", splited[0])
	assertElementInSlice(t, splited[1], freeEmailDomains)
}

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

func TestFakerMacAddress(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^[a-f0-9]{2}:[a-f0-9]{2}:[a-f0-9]{2}:[a-f0-9]{2}:[a-f0-9]{2}:[a-f0-9]{2}$", f.MacAddress())
}

func TestFakerBrowser(t *testing.T) {
	f := NewDefault()
	b := f.Browser()
	assertElementInSlice(t, b, browsers)
}
