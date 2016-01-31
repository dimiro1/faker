package defaults

import "testing"

type testStruct struct {
	AString      string  `default:"Claudemiro"`
	AInt         int     `default:"27"`
	AInt32       int32   `default:"28"`
	AInt64       int64   `default:"29"`
	AFloat32     float32 `default:"3.1415"`
	AFloat64     float64 `default:"3.1415"`
	ABoolean     bool    `default:"true"`
	NotAnnotated float64
}

type invalidTypeInStruct struct {
	InvalidType []byte `default:"invalid"`
}

type invalidTypeInStructNotAnnotated struct {
	InvalidType []byte
}

type wrongValueForFloat struct {
	AFloat64 float64 `default:"invalid"`
}

type wrongValueForInt struct {
	AInt64 int64 `default:"invalid"`
}

func TestApplyWrongValueForFloat(t *testing.T) {
	e := wrongValueForFloat{}

	err := Apply(&e)

	if err == nil {
		t.Error(err)
	}
}

func TestApplyWrongValueForInt(t *testing.T) {
	e := wrongValueForInt{}

	err := Apply(&e)

	if err == nil {
		t.Error(err)
	}
}

func TestApplyInvalidTypeInStruct(t *testing.T) {
	e := invalidTypeInStruct{}

	err := Apply(&e)

	if err == nil {
		t.Error(err)
	}
}

func TestApplyInvalidTypeInStructNotAnnotated(t *testing.T) {
	e := invalidTypeInStructNotAnnotated{}

	err := Apply(&e)

	if err != nil {
		t.Error(err)
	}
}

func TestApplyWrongValue(t *testing.T) {
	if err := Apply(1); err == nil {
		t.Error(err)
	}
}

func TestApplyDefaults(t *testing.T) {
	m := "The default value must be applied"
	e := testStruct{}

	Apply(&e)

	if e.AString != "Claudemiro" {
		t.Error(m)
		return
	}

	if e.AInt != 27 {
		t.Error(m)
		return
	}

	if e.AInt32 != 28 {
		t.Error(m)
		return
	}

	if e.AInt64 != 29 {
		t.Error(m)
		return
	}

	if e.AFloat32 != 3.1415 {
		t.Error(m)
		return
	}

	if e.AFloat64 != 3.1415 {
		t.Error(m)
		return
	}

	if !e.ABoolean {
		t.Error(m)
		return
	}

	if e.NotAnnotated != 0 {
		t.Error("This field must not be set")
		return
	}
}

func TestApplyDefaultsAvoidingAlreadySetFields(t *testing.T) {
	e := testStruct{AString: "Hello World"}

	Apply(&e)

	if e.AString != "Hello World" {
		t.Error("Must not set default, because the value is already set.")
		return
	}
}
