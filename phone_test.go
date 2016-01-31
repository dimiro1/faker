package faker

import "testing"

// How can I test this?
func TestFakerPhoneNumber(t *testing.T) {
	f := NewDefault()
	p := f.PhoneNumber()

	if len(p) == 0 {
		t.Error("Should return a phone number")
	}
}

func TestFakerPhoneAreaCode(t *testing.T) {
	f := NewDefault()
	l := f.PhoneAreaCode()
	assertElementInSlice(t, l, f.CurrentLocale().AreaCodes)
}

func TestFakerPhoneExchangeCode(t *testing.T) {
	f := NewDefault()
	l := f.PhoneExchangeCode()
	assertElementInSlice(t, l, f.CurrentLocale().ExchangeCodes)
}
