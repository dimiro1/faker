package faker

func (f Faker) BookTitle() string {
	return "BookTitle"
}

// BookAuthor returns a fake author name
func (f Faker) BookAuthor() string {
	return f.Name()
}

// BookPublisher return a publisher company name
func (f Faker) BookPublisher() string {
	return template("BookPublisher", randomElement(f.CurrentLocale().BookPublishersFormats), f)
}

// BookPublisherSuffix return a publisher company suffix
func (f Faker) BookPublisherSuffix() string {
	return randomElement(f.CurrentLocale().BookPublishersSuffixes)
}

// BookGenre return a book genre name
func (f Faker) BookGenre() string {
	return randomElement(f.CurrentLocale().BookGenres)
}

// BookISBN returns a fake ISBN number
// The ISBN number is just an EAN13 code.
// See https://pt.wikipedia.org/wiki/International_Standard_Book_Number
func (f Faker) BookISBN() string {
	return f.EAN13()
}
