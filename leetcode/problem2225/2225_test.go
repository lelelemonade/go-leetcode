package problem2225

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test2225(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 10}, {4, 5, 7, 8}}, findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
}

func findWinners(matches [][]int) [][]int {
	loseCount := make(map[int]int)
	allPlayer := make(map[int]struct{})

	for _, v := range matches {
		loseCount[v[1]]++
		allPlayer[v[0]] = struct{}{}
		allPlayer[v[1]] = struct{}{}
	}

	result := make([][]int, 2)
	result[0] = make([]int, 0)
	result[1] = make([]int, 0)

	for k, _ := range allPlayer {
		if _, e := loseCount[k]; !e {
			result[0] = append(result[0], k)
		}

		if 1 == loseCount[k] {
			result[1] = append(result[1], k)
		}
	}

	slices.Sort(result[0])
	slices.Sort(result[1])

	return result
}
