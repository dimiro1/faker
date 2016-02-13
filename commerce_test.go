package faker

import "testing"

func TestFakerProductPrice(t *testing.T) {
	f := NewDefault()
	price := f.ProductPrice()

	if price < 0 || price > 5000 {
		t.Errorf("ProductPrice must return prices from 0 inclusive to 5000 inclusive: %f", 5000)
	}
}

func TestFakerStoreDepartment(t *testing.T) {
	f := NewDefault()
	d := f.StoreDepartment()
	assertElementInSlice(t, d, f.CurrentLocale().StoreDepartments)
}

func TestFakerProductSimpleName(t *testing.T) {
	f := NewDefault()
	p := f.ProductSimpleName()
	assertElementInSlice(t, p, f.CurrentLocale().ProductNames)
}

func TestFakerProductAdjective(t *testing.T) {
	f := NewDefault()
	p := f.ProductAdjective()
	assertElementInSlice(t, p, f.CurrentLocale().ProductAdjectives)
}
