package leetcode

import (
	"fmt"
	"testing"
)

func Test151(t *testing.T) {
	fmt.Println(reverseWords("help me please"))
}

func reverseWords(s string) string {
	sArray := []byte(s)

	start := 0
	spaceBefore := false

	for i := 0; i < len(sArray); i++ {
		if sArray[i] == ' ' && spaceBefore {

		} else if sArray[i] == ' ' && !spaceBefore {
			spaceBefore = true
			sArray[start] = sArray[i]
			start++
		} else {
			spaceBefore = false
			sArray[start] = sArray[i]
			start++
		}

		if i == len(sArray)-1 {
			sArray = sArray[:start+1]
		}
	}

	reverseRange(sArray, 0, len(sArray)-1)

	end := 0
	for i := 0; i < len(sArray); {
		for sArray[end] != ' ' {
			end++
		}
		reverseRange(sArray, i, end)
		i += end
	}

	return string(sArray)
}

func reverseRange(sArray []byte, start int, end int) {
	for start < end {
		sArray[start], sArray[end] = sArray[end], sArray[start]
		start++
		end--
	}
}
