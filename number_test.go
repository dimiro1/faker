package faker

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {
	f := NewDefaultFaker()
	size := 5
	assertStringRegexp(t, fmt.Sprintf("^\\d{%d}$", size), f.Number(size))
}

func TestNumberWithWrongParam(t *testing.T) {
	f := NewDefaultFaker()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Should panic with wrong digits parameter")
		}
	}()

	f.Number(0)
}

func TestHexadecimal(t *testing.T) {
	f := NewDefaultFaker()
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

func TestNumberBetween(t *testing.T) {
	f := NewDefaultFaker()

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
