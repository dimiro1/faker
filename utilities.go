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

func randomElement(s []string) string {
	return s[rand.Intn(len(s))]
}

func randomBoolean() bool {
	return rand.Float64() >= 0.5
}

func bothify(s string) string {
	return randomReplaceAll("&", s, alpha)
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

func randomDigit() int {
	return randomIntBetween(0, 9)
}

func randomNumber(len int) int {
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

func randomFloatBetween(first, second float64) float64 {
	return rand.Float64()*float64(second-first+1) + float64(first)
}

func randomIntBetween(first, second int) int {
	return int(math.Floor(rand.Float64()*float64(second-first+1)) + float64(first))
}

func randomElements(len int) []int {
	var nums []int

	for i := 0; i < len; i++ {
		nums = append(nums, randomIntBetween(0, 9))
	}

	return nums
}

func lexify(s string) string {
	return randomReplaceAll("\\?", s, letters)
}

func hexify(s string) string {
	return randomReplaceAll("\\*", s, hex)
}

func numerify(s string) string {
	return randomReplaceAll("#", s, numbers)
}

func numerifyGreaterThanZero(s string) string {
	return randomReplaceAll("@", s, noZero)
}

func randomReplaceAll(sym, s string, slice []string) string {
	r := regexp.MustCompile(sym)

	return r.ReplaceAllStringFunc(s, func(original string) string {
		return randomElement(slice)
	})
}
