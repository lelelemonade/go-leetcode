package problem3453

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func separateSquares(squares [][]int) float64 {
    numerator := 0
	denominator := 0

	for _, square := range squares {
		numerator += (2*square[1]+square[2])*square[2]
		denominator += 2*square[2]
	}

	ans := float64(numerator)/float64(denominator)
	var minAns float64

	for _, square := range squares {
		if float64(square[1])>ans {
			continue
		}
		if float64(square[1])<ans && float64(square[1])+ float64(square[2]) > ans {
			return ans
		}
		minAns = max(minAns, float64(square[1])+ float64(square[2]))
	}

	return minAns
}

func Test3453(t *testing.T) {
	var x float64 = 12.875

	assert.Equal(t, x,separateSquares([][]int{
		{2,5,3},
		{8,12,4},
	}))
}