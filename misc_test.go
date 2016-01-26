package faker

import "testing"

func TestFakerMD5(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^[\\d0-f]{32}$", f.MD5())
}

func TestFakerSHA1(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^[\\d0-f]{40}$", f.SHA1())
}

func TestFakerSHA256(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^[\\d0-f]{64}$", f.SHA256())
}

func TestFakerSHA512(t *testing.T) {
	f := NewDefault()
	assertStringRegexp(t, "^[\\d0-f]{128}$", f.SHA512())
}
