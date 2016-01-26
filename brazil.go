package faker

import "fmt"

// BrazilCNPJ Generates a valid Brazilian CNPJ
// See: https://pt.wikipedia.org/wiki/Cadastro_Nacional_da_Pessoa_Jur%C3%ADdica
func (f Faker) BrazilCNPJ() string {
	v := [2]int{}
	cnpj := sliceOfRandonNumbers(12)
	cnpj[8] = 0
	cnpj[9] = 0
	cnpj[10] = 0

	// First Digit
	m0 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	for i, n := range m0 {
		v[0] += n * cnpj[i]
	}

	v[0] = 11 - v[0]%11

	if v[0] >= 10 {
		v[0] = 0
	}

	// Second Digit
	m1 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3}

	for i, n := range m1 {
		v[1] += n * cnpj[i]
	}

	v[1] += 2 * v[0]
	v[1] = 11 - v[1]%11

	if v[1] >= 10 {
		v[1] = 0
	}

	return fmt.Sprintf(
		"%d%d.%d%d%d.%d%d%d/%d%d%d%d-%d%d",
		cnpj[0], cnpj[1], cnpj[2], cnpj[3], cnpj[4], cnpj[5],
		cnpj[6], cnpj[7], cnpj[8], cnpj[9], cnpj[10], cnpj[11], v[0], v[1],
	)
}

// BrazilCPF Generates a valid Brazilian CPF
// See https://en.wikipedia.org/wiki/Cadastro_de_Pessoas_F%C3%ADsicas
func (f Faker) BrazilCPF() string {
	v := [2]int{}
	cpf := sliceOfRandonNumbers(11)

	// First Digit
	for i := 1; i <= 9; i++ {
		v[0] += i * cpf[i-1]
	}

	v[0] = v[0] % 11
	v[0] = v[0] % 10

	// Second Digit
	for i := 1; i <= 8; i++ {
		v[1] += i * cpf[i]
	}

	v[1] += 9 * v[0]
	v[1] = v[1] % 11
	v[1] = v[1] % 10

	return fmt.Sprintf(
		"%d%d%d.%d%d%d.%d%d%d-%d%d",
		cpf[0], cpf[1], cpf[2], cpf[3], cpf[4],
		cpf[5], cpf[6], cpf[7], cpf[8], v[0], v[1],
	)
}
