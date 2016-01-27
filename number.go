package faker

import (
	"fmt"
	"math"
)

// Decimal returns a decimal number with digits length and places decimal places length
// it will panic if len is less than 1
func (f Faker) Decimal(len, places int) string {
	return fmt.Sprintf("%d.%d", f.Number(len), f.Number(places))
}

// Digit returns an positive int between 0 (inclusive) and 9 (inclusive)
func (f Faker) Digit() int {
	return randomDigit()
}

// Hexadecimal returns a alpha numeric with digits length
// it will panic if len is less than 1
func (f Faker) Hexadecimal(len int) string {
	return hexify(fillString("*", len))
}

// NegativeNumber returns an positive int between math.MinInt32 (inclusive) and -1 (inclusive)
func (f Faker) NegativeNumber() int {
	return randomIntBetween(math.MinInt32, -1)
}

// Number returns a number with digits length
// it will panic if len is less than 1
func (f Faker) Number(len int) int {
	return randomNumber(len)
}

// NumberBetween returns a random integer between first (inclusive) and second (inclusive)
func (f Faker) NumberBetween(first, second int) int {
	return randomIntBetween(first, second)
}

// PositiveNumber returns an positive int between 0 (inclusive) and math.MaxUint32 (inclusive)
func (f Faker) PositiveNumber() int {
	return randomIntBetween(0, math.MaxUint32)
}
