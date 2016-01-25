package main

import (
	"fmt"

	"github.com/dimiro1/faker"
)

func main() {
	fake := faker.NewDefaultFaker()
	fmt.Println(fake.Decimal(4, 5))
	fmt.Println(fake.Digit())
	fmt.Println(fake.Hexadecimal(10))
	fmt.Println(fake.NegativeNumber())
	fmt.Println(fake.Number(4))
	fmt.Println(fake.NumberBetween(10, 20))
	fmt.Println(fake.PositiveNumber())
	fmt.Printf("%s %s %s\n", fake.AppName(), fake.AppVersion(), fake.AppAuthor())
}