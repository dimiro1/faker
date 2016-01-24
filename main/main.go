package main

import (
	"fmt"

	"github.com/dimiro1/faker"
)

func main() {
	fake := faker.NewDefaultFaker()
	fmt.Println(fake.City())
	fmt.Println(fake.UserAgent())
	fmt.Printf("%s %s %s\n", fake.AppName(), fake.AppVersion(), fake.AppAuthor())
}
