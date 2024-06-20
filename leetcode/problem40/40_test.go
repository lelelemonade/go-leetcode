package problem40

import (
	"fmt"
	"slices"
	"testing"
)

func Test40(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	results := make([][]int, 0)

	backtrack(candidates, []int{}, target, &results)

	return results
}

func backtrack(candidates []int, chosen []int, target int, results *[][]int) {
	if target == 0 {
		result := make([]int, len(chosen))
		copy(result, chosen)
		*results = append(*results, result)
	} else if target < 0 {
		return
	}

	for i := 0; i < len(candidates); i++ {
		backtrack(candidates[i+1:], append(chosen, candidates[i]), target-candidates[i], results)
		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
	}
}
