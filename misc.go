package faker

// Boolean returns a random true or false
func (f Faker) Boolean() bool {
	return randomBoolean()
}

// MD5 returns a random MD5 hexdigest
// See: https://en.wikipedia.org/wiki/MD5
func (f Faker) MD5() string {
	return hexify(fillString("*", 32))
}

// SHA1 returns a random SHA1 hexdigest
// See: https://en.wikipedia.org/wiki/SHA-1
func (f Faker) SHA1() string {
	return hexify(fillString("*", 40))
}

// SHA256 returns a random SHA256 hexdigest
// See: https://en.wikipedia.org/wiki/SHA-2
func (f Faker) SHA256() string {
	return hexify(fillString("*", 64))
}

// SHA512 returns a random SHA512 hexdigest
// See: https://en.wikipedia.org/wiki/SHA-2
func (f Faker) SHA512() string {
	return hexify(fillString("*", 128))
}
