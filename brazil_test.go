package faker

import "testing"

func TestBrazilCPF(t *testing.T) {
	f := NewDefaultFaker()
	assertStringRegexp(t, "\\d{3}\\.\\d{3}\\.\\d{3}-\\d{2}", f.BrazilCPF())
}

func TestBrazilCNPJ(t *testing.T) {
	f := NewDefaultFaker()
	assertStringRegexp(t, "\\d{2}\\.\\d{3}\\.\\d{3}/\\d{4}-\\d{2}", f.BrazilCNPJ())
}
