package faker

import (
	"bytes"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	gotemplate "text/template"
)

// randomElement select a random element from a array
func randomElement(s []string) string {
	return s[rand.Intn(len(s))]
}

// randomBoolean generates a random boolean
func randomBoolean() bool {
	return rand.Float64() >= 0.5
}

// bothify substitutes all & elements from alpha chars
func bothify(s string) string {
	return randomReplaceAll("&", s, alpha)
}

// fillString create a string with len size filled with sym chars
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

// randomDigit returns a random number between 0 and 9
func randomDigit() int {
	return randomIntBetween(0, 9)
}

// randomNumber returns a number with len size. Note the first char is always greater than 0
func randomNumber(len int) int {
	// The first char can not be zero
	s := "@"
	s += fillString("#", len-1)
	s = numerify(numerifyGreaterThanZero(s))

	n, err := strconv.Atoi(s)

	if err != nil {
		panic("randomNumber: This error should never happen. Please take a look at the numbers slice")
	}

	return n
}

// randomFloatBetween returns a float64 between first and second
func randomFloatBetween(first, second float64) float64 {
	return rand.Float64()*float64(second-first+1) + float64(first)
}

// randomIntBetween returns a int between first and second
func randomIntBetween(first, second int) int {
	return int(math.Floor(rand.Float64()*float64(second-first+1)) + float64(first))
}

// randomElements returns a random slice of length len
func randomElements(len int) []int {
	var nums []int

	for i := 0; i < len; i++ {
		nums = append(nums, randomIntBetween(0, 9))
	}

	return nums
}

// format apply various substitutions in the string s
func format(s string) string {
	s1 := lexify(s)
	s1 = hexify(s1)
	s1 = numerify(s1)
	s1 = bothify(s1)
	s1 = numerifyGreaterThanZero(s1)

	return s1
}

// lexify substitutes all ? elements with chars from a to z
func lexify(s string) string {
	return randomReplaceAll("\\?", s, atoz)
}

// hexify substitutes all * elements with hex chars
func hexify(s string) string {
	return randomReplaceAll("\\*", s, hex)
}

// numerify substitutes all # elements with numbers from zero to nine
func numerify(s string) string {
	return randomReplaceAll("#", s, zeroToNine)
}

// numerifyGreaterThanZero substitutes all @ elements with numbers from one to nine
func numerifyGreaterThanZero(s string) string {
	return randomReplaceAll("@", s, oneToNine)
}

// randomReplaceAll replaces all sym chars in the string s with random chars included in the slice slice
func randomReplaceAll(sym, s string, slice []string) string {
	r := regexp.MustCompile(sym)

	return r.ReplaceAllStringFunc(s, func(original string) string {
		return randomElement(slice)
	})
}

func template(str string, data interface{}) (string, error) {
	buffer := new(bytes.Buffer)

	t := gotemplate.Must(gotemplate.New("template").Parse(str))
	err := t.Execute(buffer, data)

	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
