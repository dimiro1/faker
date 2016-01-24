package faker

func (f Faker) EAN() string {
	return "EAN"
}

func (f Faker) Number(digits int) string {
	return "1968353479"
}

func (f Faker) Decimal(digits, places int) string {
	return "19683.53479"
}

func (f Faker) Hexadecimal(digits int) string {
	return "a12f4cb"
}

func (f Faker) NumberBetween(n1, n2 int) int {
	return 7
}

func (f Faker) PositiveNumber() int {
	return 28
}

func (f Faker) NegativeNumber() int {
	return -28
}
