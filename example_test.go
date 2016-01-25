package faker_test

import (
	"fmt"

	"github.com/dimiro1/faker"
)

func ExampleFaker_BrazilCPF() {
	f := faker.NewDefaultFaker()
	fmt.Println(f.BrazilCPF())
	// Output: 251.448.621-10
}
