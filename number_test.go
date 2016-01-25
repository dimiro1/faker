package faker

import (
	"fmt"
	"strconv"
	"testing"
)

var f Faker

func init() {
	f = NewDefaultFaker()
}

func TestDecimal(t *testing.T) {
	digits := 5
	places := 5
	assertStringRegexp(t, fmt.Sprintf("^\\d{%d}\\.\\d{%d}$", digits, places), f.Decimal(digits, places))
}

func TestDigit(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := f.Digit()
		if n < 0 || n > 9 {
			t.Errorf("Digit must return digits from 0 inclusive to 9 inclusive: %d", n)
		}
	}
}

func TestHexadecimal(t *testing.T) {
	size := 5
	assertStringRegexp(t, fmt.Sprintf("^[\\da-f]{%d}$", size), f.Hexadecimal(size))
}

func TestHexadecimalWithWrongParam(t *testing.T) {
	f := NewDefaultFaker()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Should panic with wrong digits parameter")
		}
	}()

	f.Hexadecimal(0)
}

func TestNegativeNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := f.NegativeNumber()
		if n >= 0 {
			t.Errorf("NegativeNumber must return only negative numbers: %d", n)
		}
	}
}

func TestNumber(t *testing.T) {
	size := 5
	n := strconv.Itoa(f.Number(size))
	assertStringRegexp(t, fmt.Sprintf("^\\d{%d}$", size), n)
}

func TestNumberWithWrongParam(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Should panic with wrong digits parameter")
		}
	}()

	f.Number(0)
}

func TestNumberBetween(t *testing.T) {
	input := [][]int{
		{0, 0},
		{0, 10},
		{-10, -1},
		{-5, 5},
		{10, 0},
		{-1, -10},
	}

	for _, e := range input {
		first := e[0]
		second := e[1]
		n := f.NumberBetween(first, second)

		if !((n >= first && n <= second) || (n <= first && n >= second)) {
			t.Errorf("Number must be between first inclusive and second inclusive: (n=%d, first=%d, second=%d)", n, first, second)
		}
	}
}

func TestPositiveNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := f.PositiveNumber()
		if n < 0 {
			t.Errorf("PositiveNumber must return only positive numbers: %d", n)
		}
	}
}
