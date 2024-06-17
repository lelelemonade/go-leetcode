package problem39

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test39(t *testing.T) {
	assert.Equal(t, [][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 46, 7}, 7))
	assert.Equal(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8))
}

func TestSlices(t *testing.T) {
	a := []int{1}
	b := []int{2}

	MyTest(a, &b)

	fmt.Println(a)
	fmt.Println(b)
}

func MyTest(a []int, b *[]int) {
	fmt.Println(a)
	fmt.Println(*b)
	a = append(a, 3)
	*b = append(*b, 4)
	fmt.Println(a)
	fmt.Println(*b)
}

func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	slices.Sort(candidates)

	backtrack(candidates, target, 0, []int{}, &results)

	return results
}

func backtrack(candidates []int, target int, startIdx int, result []int, results *[][]int) {
	if target == 0 {
		resultCopy := make([]int, len(result))
		copy(resultCopy, result)
		*results = append(*results, resultCopy)
	} else if target < 0 {
		return
	}

	for i := startIdx; i < len(candidates); i++ {
		backtrack(candidates, target-candidates[i], i, append(result, candidates[i]), results)
	}
}
