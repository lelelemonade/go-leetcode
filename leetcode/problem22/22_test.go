package problem22

import (
	"fmt"
	"testing"
)

func Test22(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	backtrack("", n, n, &result)
	return result
}

func backtrack(candidate string, left, right int, result *[]string) {
	if left == 0 {
		rightArray := make([]byte, right)
		for i := 0; i < len(rightArray); i++ {
			rightArray[i] = ')'
		}
		*result = append(*result, candidate+string(rightArray))
		return
	}

	if left < right {
		backtrack(candidate+"(", left-1, right, result)
		backtrack(candidate+")", left, right-1, result)
	} else if left == right {
		backtrack(candidate+"(", left-1, right, result)
	}
}
