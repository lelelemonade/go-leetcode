package problem17

import (
	"fmt"
	"testing"
)

var numberArray = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func Test17(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var result []string
	var numberToIter [][]byte
	for i := 0; i < len(digits); i++ {
		numberToIter = append(numberToIter, numberArray[digits[i]])
	}

	backtrack(numberToIter, []byte{}, &result)

	return result
}

func backtrack(numberToIter [][]byte, alreadyAdded []byte, result *[]string) {
	if len(numberToIter) == 0 {
		*result = append(*result, string(alreadyAdded))
		return
	}

	for i := 0; i < len(numberToIter[0]); i++ {
		backtrack(numberToIter[1:], append(alreadyAdded, numberToIter[0][i]), result)
	}
}
