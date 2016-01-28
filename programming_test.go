package faker

import "testing"

func TestFakerProgrammingLanguage(t *testing.T) {
	f := NewDefault()
	l := f.ProgrammingLanguage()
	assertElementInSlice(t, l, programmingLanguages)
}

func TestFakerOS(t *testing.T) {
	f := NewDefault()
	l := f.OS()
	assertElementInSlice(t, l, operatingSystems)
}

func TestFakerTextEditor(t *testing.T) {
	f := NewDefault()
	l := f.TextEditor()
	assertElementInSlice(t, l, textEditors)
}
