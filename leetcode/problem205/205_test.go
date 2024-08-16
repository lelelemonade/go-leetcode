package problem205

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isIsomorphic(s string, t string) bool {
	sIdx := make([]int, 128)
	tIdx := make([]int, 128)

	for i := 0; i < len(s); i++ {
		if sIdx[s[i]] == 0 && tIdx[t[i]] == 0 {
			sIdx[s[i]] = i + 1
			tIdx[t[i]] = i + 1
		} else if sIdx[s[i]] != 0 && tIdx[t[i]] != 0 {
			if sIdx[s[i]] != tIdx[t[i]] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func Test205(t *testing.T) {
	assert.Equal(t, false, isIsomorphic("ab", "aa"))
}
