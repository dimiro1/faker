package faker

import "image/color"

func (f Faker) HexColor() string {
	return "HexColor"
}

func (f Faker) ColorName() string {
	return "ColorName"
}

func (f Faker) RGBA() color.RGBA {
	return color.RGBA{R: 0, G: 0, B: 0, A: 1}
}

func (f Faker) CMYK() color.CMYK {
	return color.CMYK{C: 0, M: 0, Y: 0, K: 0}
}
