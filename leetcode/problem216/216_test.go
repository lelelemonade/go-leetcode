package problem216

import (
	"fmt"
	"testing"
)

func Test216(t *testing.T) {
	result := combine(7, 3)
	result = combine(9, 3)

	fmt.Println(result)
}

func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	backtrack(0, n, k, map[int]struct{}{}, &result)
	return result
}

func backtrack(right int, count int, target int, alreadyCountedInt map[int]struct{}, result *[][]int) {
	if len(alreadyCountedInt) == target-1 {
		for i := right + 1; i <= 9; i++ {
			currentCount := 0
			var resultInts []int
			for k := range alreadyCountedInt {
				currentCount += k
				resultInts = append(resultInts, k)
			}

			if currentCount != count-i {
				continue
			}

			resultInts = append(resultInts, i)
			*result = append(*result, resultInts)
		}
		return
	}

	for i := right + 1; i < 10; i++ {
		nextCountedInt := make(map[int]struct{}, len(alreadyCountedInt)+1)
		for k := range alreadyCountedInt {
			nextCountedInt[k] = struct{}{}
		}
		nextCountedInt[i] = struct{}{}
		backtrack(i, count, target, nextCountedInt, result)
	}
}
