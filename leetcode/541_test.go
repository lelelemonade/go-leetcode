package leetcode

import (
	"fmt"
	"testing"
)

func Test541(t *testing.T) {
	fmt.Println(reverseStr("abcdefg", 2))
}

func reverseStr(s string, k int) string {
	sBytes := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		start := i
		end := min(i+k-1, len(s)-1)
		for start < end {
			sBytes[start], sBytes[end] = sBytes[end], sBytes[start]
			start++
			end--
		}
	}
	return string(sBytes)
}
