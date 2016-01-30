package faker

// OS returns a name of an Operating System
func (f Faker) OS() string {
	return randomElement(operatingSystems)
}

// ProgrammingLanguage returns a name of a programming language
func (f Faker) ProgrammingLanguage() string {
	return randomElement(programmingLanguages)
}

// TextEditor returns a name of a TextEditor
func (f Faker) TextEditor() string {
	return randomElement(textEditors)
}

// WebFramework returns a name of a WebFramework
func (f Faker) WebFramework() string {
	return randomElement(webFrameworks)
}
