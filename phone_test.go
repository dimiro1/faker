package faker

import "testing"

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
