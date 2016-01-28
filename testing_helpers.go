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

func assertElementInSlice(t *testing.T, s string, slice []string) {
	found := false
	for _, c := range slice {
		if c == s {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("%s Is present on the list", s)
	}
}
