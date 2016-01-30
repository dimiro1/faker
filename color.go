package faker

import "image/color"

// ColorName returns a name of a ColorName
func (f Faker) ColorName() string {
	return randomElement(colorNames)
}

// CMYK returns a valid CMYK color
func (f Faker) CMYK() color.CMYK {
	r := uint8(randomIntBetween(0, 255))
	g := uint8(randomIntBetween(0, 255))
	b := uint8(randomIntBetween(0, 255))

	c, m, y, k := color.RGBToCMYK(r, g, b)

	return color.CMYK{C: c, M: m, Y: y, K: k}
}

// HexColor returns a hexadecimal color, including alpha
func (f Faker) HexColor() string {
	return hexify(fillString("*", 8))
}

// RGBA returns a valid RGBA color
func (f Faker) RGBA() color.RGBA {
	r := uint8(randomIntBetween(0, 255))
	g := uint8(randomIntBetween(0, 255))
	b := uint8(randomIntBetween(0, 255))
	a := uint8(randomIntBetween(0, 255))

	return color.RGBA{R: r, G: g, B: b, A: a}
}
