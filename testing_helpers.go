package faker

import (
	"regexp"
	"testing"
)

// Help test regular expressions
func assertStringRegexp(t *testing.T, pattern, s string) {
	matches, _ := regexp.MatchString(pattern, s)

	if !matches {
		t.Errorf("%s is not in format %s", s, pattern)
	}
}
