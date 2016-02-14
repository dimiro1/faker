package faker

import "testing"

func TestFakerCompanyLogo(t *testing.T) {
	f := NewDefault()
	l := f.CompanyLogo()
	assertElementInSlice(t, l, companyLogos)
}
