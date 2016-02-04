package faker

import "testing"

func TestBothify(t *testing.T) {
	s := bothify("&&&")
	assertStringRegexp(t, "^\\w{3}$", s)
}

func TestInternalDigit(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := randomDigit()
		if n < 0 || n > 9 {
			t.Errorf("Digit must return digits from 0 inclusive to 9 inclusive: %d", n)
		}
	}
}

func TestFillString(t *testing.T) {
	s := fillString("#", 10)
	assertStringRegexp(t, "^[#]{10}$", s)
}

func TestFormat(t *testing.T) {
	s := format("#@*?")
	assertStringRegexp(t, "^\\d[1-9][\\da-f][a-z]$", s)
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

func TestTemplate(t *testing.T) {
	s, err := template("test-template", `Hello {{ .Name }}`, struct {
		Name string
	}{Name: "Claudemiro"})

	if err != nil {
		t.Error("Should not return error")
	}

	if s != "Hello Claudemiro" {
		t.Errorf("The result is different '%s'", s)
	}

}
