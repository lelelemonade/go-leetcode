package problem2730

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestSemiRepetitiveSubstring(s string) int {
	if len(s)==1{
		return 1
	}
    semiRepCount:=0
	left:=0
	longest:=0
	for right:=1;right<len(s);right++{
		if s[right]==s[right-1]{
			semiRepCount++
		}
		for semiRepCount>1{
			if s[left]==s[left+1]{
				semiRepCount--
			}
			left++
		}
		longest=max(longest,right-left+1)
	}
	return longest
}

func Test2730(t *testing.T) {
	assert.Equal(t,4,longestSemiRepetitiveSubstring("52233"))
}