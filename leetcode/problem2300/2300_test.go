package problem2300

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
    res := make([]int, len(spells))
    for i, x := range spells {
        res[i] = len(potions) - sort.SearchInts(potions, (int(success) - 1) / x + 1)
    }
    return res
}

func Test2300(t *testing.T) {
	// assert.Equal(t,[]int{2,0,2},successfulPairs([]int{3,1,2},[]int{8,5,8},16))
	assert.Equal(t,[]int{3,0,3},successfulPairs([]int{15,8,19},[]int{38,36,23},328))
}