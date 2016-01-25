package faker

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// Decimal returns a decimal number with digits length and places decimal places length
// it will panic if digits is less than 1
func (f Faker) Decimal(digits, places int) string {
	return fmt.Sprintf("%d.%d", f.Number(digits), f.Number(places))
}

// Digit returns an positive int between 0 (inclusive) and 9 (inclusive)
func (f Faker) Digit() int {
	return f.NumberBetween(0, 9)
}

// Hexadecimal returns a alpha numeric with digits length
// it will panic if digits is less than 1
func (f Faker) Hexadecimal(digits int) string {
	return hexify(fillString("*", digits))
}

// NegativeNumber returns an positive int between math.MinInt32 (inclusive) and -1 (inclusive)
func (f Faker) NegativeNumber() int {
	return f.NumberBetween(math.MinInt32, -1)
}

// Number returns a number with digits length
// it will panic if digits is less than 1
// TODO: Remove usage of randomElements in favor of numerify, lexify etc functions
func (f Faker) Number(digits int) int {
	// The first element can not be zero
	s := anyFromSlice(noZero)
	s += randomElements(digits-1, numbers, "Number")

	n, err := strconv.Atoi(s)

	if err != nil {
		panic("Number: This error should never happen. Please take a look at the numbers slice")
	}

	return n
}

// NumberBetween returns a random integer between first (inclusive) and second (inclusive)
func (f Faker) NumberBetween(first, second int) int {
	return int(math.Floor(rand.Float64()*float64(second-first+1)) + float64(first))
}

// PositiveNumber returns an positive int between 0 (inclusive) and math.MaxUint32 (inclusive)
func (f Faker) PositiveNumber() int {
	return f.NumberBetween(0, math.MaxUint32)
}
