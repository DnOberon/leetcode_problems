package palindrome

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestPalindrome(t *testing.T) {
	result := longestPalindrome("babad")

	assert.Equal(t, result, "aba")
	t.Log(result)
}

func TestPalindrome2(t *testing.T) {
	result := longestPalindrome("cbbd")

	assert.Equal(t, result, "bb")
	t.Log(result)
}

func TestPalindrome3(t *testing.T) {
	result := longestPalindrome("bb")

	assert.Equal(t, result, "bb")
	t.Log(result)
}

func TestPalindrome4(t *testing.T) {
	result := longestPalindrome("ccc")

	assert.Equal(t, result, "ccc")
	t.Log(result)
}

func TestPalindrome5(t *testing.T) {
	result := longestPalindrome("abaxabaxabb")

	assert.Equal(t, result, "baxabaxab")
	t.Log(result)
}
