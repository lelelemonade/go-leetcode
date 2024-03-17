package problem77

import (
	"fmt"
	"testing"
)

func Test77(t *testing.T) {
	result := combine(4, 2)
	fmt.Println(result)
}

func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	backtrack(0, n, k, map[int]struct{}{}, &result)
	return result
}

func backtrack(right int, end int, target int, alreadyCountedInt map[int]struct{}, result *[][]int) {
	if len(alreadyCountedInt) == target-1 {
		for i := right + 1; i <= end; i++ {
			var resultInts []int
			for k := range alreadyCountedInt {
				resultInts = append(resultInts, k)
			}
			resultInts = append(resultInts, i)
			*result = append(*result, resultInts)
		}
		return
	}

	for i := right + 1; i < end; i++ {
		nextCountedInt := make(map[int]struct{}, len(alreadyCountedInt)+1)
		for k := range alreadyCountedInt {
			nextCountedInt[k] = struct{}{}
		}
		nextCountedInt[i] = struct{}{}
		backtrack(i, end, target, nextCountedInt, result)
	}
}
