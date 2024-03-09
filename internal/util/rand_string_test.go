package util

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStringLength(t *testing.T) {
	const n = 16
	s := RandString(n)
	assert.Equal(t, len(s), n, "RandString(%d) = %q, len = %d", n, s, len(s))

}

func TestRandStringContents(t *testing.T) {
	const n = 16
	s := RandString(n)

	matched, err := regexp.Match("^[0-9a-zA-Z]{16}$", []byte(s))

	assert.Nil(t, err, "RandString(%d) = %q, len = %d", n, s, len(s))
	assert.True(t, matched, "RandString(%d) = %q, len = %d", n, s, len(s))
}
