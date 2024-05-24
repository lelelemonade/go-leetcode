package problem1122

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test1122(t *testing.T) {
	assert.Equal(t, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}, relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	assert.Equal(t, []int{22, 28, 8, 6, 17, 44}, relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank := map[int]int{}
	for i, v := range arr2 {
		rank[v] = i - len(arr2)
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		if r, has := rank[x]; has {
			x = r
		}
		if r, has := rank[y]; has {
			y = r
		}
		return x < y
	})
	return arr1
}
