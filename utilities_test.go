package faker

import "testing"

func TestBothify(t *testing.T) {
	s := bothify("&&&")
	assertStringRegexp(t, "^\\w{3}$", s)
}

func TestFillString(t *testing.T) {
	s := fillString("#", 10)
	assertStringRegexp(t, "^[#]{10}$", s)
}

func TestNumerify(t *testing.T) {
	s := numerify("###")
	assertStringRegexp(t, "^\\d{3}$", s)
}

func TestNumerifyGreaterThanZero(t *testing.T) {
	s := numerifyGreaterThanZero("@@@")
	assertStringRegexp(t, "^[1-9]{3}$", s)
}

func TestLexify(t *testing.T) {
	s := lexify("???")
	assertStringRegexp(t, "^[a-z]{3}$", s)
}

func TestHexify(t *testing.T) {
	s := hexify("***")
	assertStringRegexp(t, "^[\\da-f]{3}$", s)
}
