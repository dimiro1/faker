package faker

import (
	"testing"
	"time"
)

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

func TestFakerCreditCardExpiryDate(t *testing.T) {
	f := NewDefault()
	c := f.CreditCardExpiryDate()

	if !c.After(time.Now()) {
		t.Error("CreditCardExpiryDate must be in the future")
	}
}
