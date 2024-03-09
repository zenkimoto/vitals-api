package util

import (
	"regexp"
	"testing"
)

func TestRandStringLength(t *testing.T) {
	const n = 16
	s := RandString(n)
	if len(s) != n {
		t.Errorf("RandString(%d) = %q, len = %d", n, s, len(s))
	}
}

func TestRandStringContents(t *testing.T) {
	const n = 16
	s := RandString(n)

	matched, err := regexp.Match("^[0-9a-zA-Z]{16}$", []byte(s))

	if !matched || err != nil {
		t.Errorf("RandString(%d) = %q, len = % d", n, s, len(s))
	}
}
