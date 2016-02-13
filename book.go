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

func (f Faker) BookGenre() string {
	return "BookGenre"
}

func (f Faker) BookISBN() string {
	return "BookISBN"
}
