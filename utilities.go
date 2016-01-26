package faker

import (
	"bytes"
	"math"
	"math/rand"
	"regexp"
	"strconv"
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

func digit() int {
	return numberBetween(0, 9)
}

func number(len int) int {
	// The first char can not be zero
	s := "@"
	s += fillString("#", len-1)
	s = numerify(numerifyGreaterThanZero(s))

	n, err := strconv.Atoi(s)

	if err != nil {
		panic("Number: This error should never happen. Please take a look at the numbers slice")
	}

	return n
}

func numberBetween(first, second int) int {
	return int(math.Floor(rand.Float64()*float64(second-first+1)) + float64(first))
}

func sliceOfRandonNumbers(len int) []int {
	var nums []int

	for i := 0; i < len; i++ {
		nums = append(nums, numberBetween(0, 9))
	}

	return nums
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

func numerifyGreaterThanZero(s string) string {
	return replace("@", s, noZero)
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
