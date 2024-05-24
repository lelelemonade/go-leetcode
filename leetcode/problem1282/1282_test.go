package problem1282

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1282(t *testing.T) {
	assert.Equal(t, [][]int{{}}, groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}

func groupThePeople(groupSizes []int) [][]int {
	groupMap := make(map[int]*[]int)

	for i, size := range groupSizes {
		if v, e := groupMap[size]; e {
			*v = append(*v, i)
		} else {
			value := make([]int, 0)
			value = append(value, i)
			groupMap[size] = &value
		}
	}

	result := make([][]int, 0)

	for k, v := range groupMap {
		for i := 0; i < len(*v); i += k {
			sizeList := make([]int, k)

			for j := 0; j < k; j++ {
				sizeList[j] = (*v)[i+j]
			}

			result = append(result, sizeList)
		}
	}

	return result
}
