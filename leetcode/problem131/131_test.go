package problem131

import (
	"fmt"
	"testing"
)

func Test131(t *testing.T) {
	//fmt.Println(partition("aabbbbbc"))
	fmt.Println(partition("cbbbcc"))
	//fmt.Println(partition("a"))
	//fmt.Println(partition("abb"))
	//fmt.Println(partition("aaa"))
	//fmt.Println(partition("abbc"))
}

func partition(s string) [][]string {
	result := [][]string{}

	backtrack(s, []byte{}, []string{}, &result)

	return result
}

func backtrack(toPartition string, potentialPalindrome []byte, existedPalindrome []string, result *[][]string) {
	if len(toPartition) == 0 && len(potentialPalindrome) == 0 {
		appendedResult := make([]string, len(existedPalindrome))
		copy(appendedResult, existedPalindrome)
		*result = append(*result, appendedResult)
		return
	}

	if len(toPartition) == 0 && len(potentialPalindrome) != 0 {
		return
	}

	thisByte := toPartition[0]

	nextPotentialPalindrome := append(potentialPalindrome, thisByte)
	backtrack(toPartition[1:], nextPotentialPalindrome, existedPalindrome, result)

	if isPalindrome(potentialPalindrome, thisByte) {
		nextExistedPalindrome := append(existedPalindrome, string(nextPotentialPalindrome))
		backtrack(toPartition[1:], []byte{}, nextExistedPalindrome, result)
	}
}

func isPalindrome(potentialPalindrome []byte, addedByte byte) bool {
	if len(potentialPalindrome) != 0 && addedByte != potentialPalindrome[0] {
		return false
	}

	for j := 1; j < (len(potentialPalindrome)+1)/2; j++ {
		if potentialPalindrome[j] != potentialPalindrome[len(potentialPalindrome)-j] {
			return false
		}
	}
	return true
}
