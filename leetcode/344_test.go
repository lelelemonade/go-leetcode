package leetcode

import (
	"fmt"
	"testing"
)

func Test344(t *testing.T) {
	testCase := []byte("A man, A plan, A canal: Panama")
	reverseString(testCase)
	fmt.Println(string(testCase))
}

func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
