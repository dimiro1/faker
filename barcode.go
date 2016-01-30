package faker

import (
	"bytes"
	"fmt"
)

// EAN13 Generates a valid EAN13 bar code
func (f Faker) EAN13() string {
	v := 0
	ean13 := randomElements(12)

	for i, n := range ean13 {
		v = v + n*((i%2)*2+1)
	}

	v = (10 - (v % 10)) % 10

	var buffer bytes.Buffer

	for _, n := range ean13 {
		buffer.WriteString(fmt.Sprintf("%d", n))
	}

	return fmt.Sprintf("%s%d", buffer.String(), v)
}
