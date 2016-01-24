package faker

func (f Faker) LoremWord() string {
	return "Lorem"
}

func (f Faker) LoremWords(l int) []string {
	return []string{"Lorem", "Ipsun"}
}

func (f Faker) LoremSentense() string {
	return "Lorem"
}

func (f Faker) LoremSentenses(l int) []string {
	return []string{"Lorem", "Ipsun"}
}

func (f Faker) LoremParagraph() string {
	return "Lorem"
}

func (f Faker) LoremParagraphs(l int) []string {
	return []string{"Lorem", "Ipsun"}
}
