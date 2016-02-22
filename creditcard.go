package faker

import "time"

// CreditCardNumber returns a valid credit card number
// for a while I will not generate a full random valid card
// I will just return from a list of valid numbers
func (f Faker) CreditCardNumber() string {
	return randomElement(crediCardNumbers)
}

// CreditCardExpiryDate returns datetime after now
func (f Faker) CreditCardExpiryDate() time.Time {
	return f.DateTimeAfter(time.Now())
}

// CreditCardType returns a valid credit card type
func (f Faker) CreditCardType() string {
	return randomElement(crediCardTypes)
}
