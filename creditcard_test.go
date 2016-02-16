package faker

import "testing"

func TestFakerCreditCardType(t *testing.T) {
	f := NewDefault()
	c := f.CreditCardType()
	assertElementInSlice(t, c, crediCardTypes)
}

func TestFakerCreditCardNumber(t *testing.T) {
	f := NewDefault()
	c := f.CreditCardNumber()
	assertElementInSlice(t, c, crediCardNumbers)
}
