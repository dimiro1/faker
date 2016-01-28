package faker

// ProgrammingLanguage returns a name of a programming language
func (f Faker) ProgrammingLanguage() string {
	return randomElement(programmingLanguages)
}

// OS returns a name of an Operating System
func (f Faker) OS() string {
	return randomElement(operatingSystems)
}

// TextEditor returns a name of a TextEditor
func (f Faker) TextEditor() string {
	return randomElement(textEditors)
}

func (f Faker) WebFramework() string {
	return "Rails"
}
