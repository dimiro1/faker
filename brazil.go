package faker

import (
	"fmt"
	"strconv"
	"strings"
)

// BrazilCNPJ Generates a valid Brazilian CNPJ
func (f Faker) BrazilCNPJ() string {
	return "45.294.670/0001-04"
}

// BrazilCPF Generates a valid Brazilian CPF
// CPF is a number composed of 11 digits
func (f Faker) BrazilCPF() string {
	cpf := "###.###.###"
	cpf = numerify(cpf)

	s := strings.Replace(cpf, ".", "", -1)

	v := [2]int{}
	sliceCpf := strings.Split(s, "")
	sliceIntCpf := [11]int{}

	for i, d := range sliceCpf {
		sliceIntCpf[i], _ = strconv.Atoi(d)
	}

	v[0] = sliceIntCpf[0] + 2*sliceIntCpf[1] + 3*sliceIntCpf[2]
	v[0] += 4*sliceIntCpf[3] + 5*sliceIntCpf[4] + 6*sliceIntCpf[5]
	v[0] += 7*sliceIntCpf[6] + 8*sliceIntCpf[7] + 9*sliceIntCpf[8]
	v[0] = v[0] % 11
	v[0] = v[0] % 10

	v[1] = sliceIntCpf[1] + 2*sliceIntCpf[2] + 3*sliceIntCpf[3]
	v[1] += 4*sliceIntCpf[4] + 5*sliceIntCpf[5] + 6*sliceIntCpf[6]
	v[1] += 7*sliceIntCpf[7] + 8*sliceIntCpf[8] + 9*v[0]
	v[1] = v[1] % 11
	v[1] = v[1] % 10

	return fmt.Sprintf("%s-%d%d", cpf, v[0], v[1])
}
