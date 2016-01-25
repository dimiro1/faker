package faker

import "testing"

func TestNumerify(t *testing.T) {
	s := numerify("Hello ###")
	assertStringRegexp(t, "Hello \\d{3}", s)
}

func TestLexify(t *testing.T) {
	s := lexify("Hello ???")
	assertStringRegexp(t, "Hello [a-z]{3}", s)
}

func TestBothify(t *testing.T) {
	s := bothify("Hello &&&")
	assertStringRegexp(t, "Hello \\w{3}", s)
}

func TestHexify(t *testing.T) {
	s := hexify("Hello ***")
	assertStringRegexp(t, "Hello [\\da-f]{3}", s)
}
