package problem1750

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		for left < right-1 && s[left] == s[left+1] {
			left++
		}
		for left < right-1 && s[right] == s[right-1] {
			right--
		}
		left++
		right--
	}
	return right - left + 1
}

func Test1750(t *testing.T) {
	assert.Equal(t, 0, minimumLength("cabaabac"))
	assert.Equal(t, 3, minimumLength("aabccabba"))
}
