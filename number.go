package faker

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
)

var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var alpha = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

func randomElementFromSlice(s []string) string {
	n := rand.Intn(len(s))

	return s[n]
}

func randomElements(digits int, slice []string, fn string) string {
	if digits < 1 {
		panic(fmt.Sprintf("%s: digits must be greater than 0", fn))
	}

	var buffer bytes.Buffer

	for i := 0; i < digits; i++ {
		buffer.WriteString(randomElementFromSlice(slice))
	}

	return buffer.String()
}

func (f Faker) EAN() string {
	return "EAN"
}

// Number returns a number with digits length
// it will panic if digits is less than 1
func (f Faker) Number(digits int) string {
	return randomElements(digits, numbers, "Number")
}

func (f Faker) Decimal(digits, places int) string {
	return "19683.53479"
}

// Hexadecimal returns a alpha numeric with digits length
// it will panic if digits is less than 1
func (f Faker) Hexadecimal(digits int) string {
	return randomElements(digits, alpha, "Hexadecimal")
}

// NumberBetween returns a random integer between first (inclusive) and second (inclusive)
func (f Faker) NumberBetween(first, second int) int {
	return int(math.Floor(rand.Float64()*float64(second-first+1)) + float64(first))
}

func (f Faker) PositiveNumber() int {
	return 28
}

func (f Faker) NegativeNumber() int {
	return -28
}
