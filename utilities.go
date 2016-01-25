package faker

import (
	"bytes"
	"fmt"
	"math/rand"
	"regexp"
)

var (
	zero       = []string{"0"}
	noZero     = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	hexletters = []string{"a", "b", "c", "d", "e", "f"}
	letters    = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z"}
	numbers = append(zero, noZero...)
	hex     = append(numbers, hexletters...)
	alpha   = append(numbers, letters...)
)

func bothify(s string) string {
	return replace("&", s, alpha)
}

func fillString(sym string, len int) string {
	if len < 1 {
		panic("fillString: digits must be greater than 0")
	}

	var buffer bytes.Buffer

	for i := 0; i < len; i++ {
		buffer.WriteString(sym)
	}

	return buffer.String()
}

func lexify(s string) string {
	return replace("\\?", s, letters)
}

func hexify(s string) string {
	return replace("\\*", s, hex)
}

func numerify(s string) string {
	return replace("#", s, numbers)
}

func randomElements(digits int, slice []string, fn string) string {
	if digits < 1 {
		panic(fmt.Sprintf("%s: digits must be greater than 0", fn))
	}

	var buffer bytes.Buffer

	for i := 0; i < digits; i++ {
		buffer.WriteString(anyFromSlice(slice))
	}

	return buffer.String()
}

func anyFromSlice(s []string) string {
	return s[rand.Intn(len(s))]
}

func replace(sym, s string, slice []string) string {
	r := regexp.MustCompile(sym)

	return r.ReplaceAllStringFunc(s, func(original string) string {
		return anyFromSlice(slice)
	})
}
