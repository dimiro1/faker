package faker

// BrazilCNPJ Generates a valid Brazilian CNPJ
func (f Faker) BrazilCNPJ() string {
	return "45.294.670/0001-04"
}

// BrazilCPF Generates a valid Brazilian CPF
// CPF is a number composed of 11 digits
func (f Faker) BrazilCPF() string {
	return "251.448.621-10"
}
