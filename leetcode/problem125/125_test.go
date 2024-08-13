package problem125

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(s string) bool {
	s=strings.ToLower(s)
	for left,right:=0,len(s)-1;left<right; {
		for !isValidAlphaNumeric(s[left]){
			left++
			if left >len(s)-1{
				return true
			}
		}
		for !isValidAlphaNumeric(s[right]){
			right--
			if right < 0 {
				return true
			}
		}

		if s[left]==s[right] {
			left++
			right--
			continue
		}
		return false
	}

	return true
}

func isValidAlphaNumeric(r byte) bool {
	return (r>='0'&&r<='9')||(r>='a'&&r<='z')||(r>='A'&&r<='Z')
}

func Test125(t *testing.T) {
	assert.Equal(t,false,isPalindrome("0P"))
}