package faker

import "testing"

func TestFakerOS(t *testing.T) {
	f := NewDefault()
	l := f.OS()
	assertElementInSlice(t, l, operatingSystems)
}

func TestFakerProgrammingLanguage(t *testing.T) {
	f := NewDefault()
	l := f.ProgrammingLanguage()
	assertElementInSlice(t, l, programmingLanguages)
}

func TestFakerTextEditor(t *testing.T) {
	f := NewDefault()
	l := f.TextEditor()
	assertElementInSlice(t, l, textEditors)
}

func TestFakerWebFramework(t *testing.T) {
	f := NewDefault()
	l := f.WebFramework()
	assertElementInSlice(t, l, webFrameworks)
}
