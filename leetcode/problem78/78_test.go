package problem78

import (
	"fmt"
	"testing"
)

func Test78(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 2}))
}

func subsets(nums []int) [][]int {
	result := [][]int{}

	notTriedSet := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		notTriedSet[num] = struct{}{}
	}

	backtrack([]int{}, notTriedSet, &result)

	return result
}

func backtrack(alreadyTriedSet []int, notTriedSet map[int]struct{}, result *[][]int) {
	resultToAppend := make([]int, len(alreadyTriedSet))

	for key, value := range alreadyTriedSet {
		resultToAppend[key] = value
	}
	*result = append(*result, resultToAppend)

	for k := range notTriedSet {
		if len(alreadyTriedSet) == 0 || (len(alreadyTriedSet) > 0 && k > alreadyTriedSet[len(alreadyTriedSet)-1]) {
			nextTrySet := append(alreadyTriedSet, k)
			nextNotTrySet := make(map[int]struct{}, len(notTriedSet))
			for key, value := range notTriedSet {
				nextNotTrySet[key] = value
			}

			delete(nextNotTrySet, k)

			backtrack(nextTrySet, notTriedSet, result)
		}
	}
}
