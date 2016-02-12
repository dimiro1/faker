package main

import (
	"fmt"

	"github.com/dimiro1/faker"
)

func main() {
	fake := faker.NewDefault()

	fmt.Println(fake.Decimal(4, 5))
	fmt.Println(fake.Digit())
	fmt.Println(fake.Hexadecimal(10))
	fmt.Println(fake.NegativeNumber())
	fmt.Println(fake.Number(4))
	fmt.Println(fake.NumberBetween(10, 20))
	fmt.Println(fake.PositiveNumber())
	fmt.Println(fake.Country())
	fmt.Println(fake.State())
	fmt.Println(fake.StateAbbr())

	fmt.Println(fake.ProgrammingLanguage())
	fmt.Println(fake.OS())
	fmt.Println(fake.EAN13())
	fmt.Println(fake.PhoneNumber())
	fmt.Println(fake.Name())
	fmt.Println(fake.CompanyName())
	fmt.Println(fake.StreetName())
	fmt.Println(fake.StreetAddress())
	fmt.Println(fake.City())

	fmt.Println(fake.Address())

	// fmt.Printf("%s %s %s\n", fake.AppName(), fake.AppVersion(), fake.AppAuthor())
}
