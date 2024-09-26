package problem3090

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximumLengthSubstring(s string) int {
	charCount := make(map[rune]int)
	left:=0
	result := 0
	for right,v:=range s{
		charCount[v]++
		for charCount[v]>2{
			charCount[rune(s[left])]--
			left++
		}
		result=max(result,right-left+1)
	}
	return result
}

func Test3090(t *testing.T) {
	assert.Equal(t,4,maximumLengthSubstring("bcbbbcba"))
}