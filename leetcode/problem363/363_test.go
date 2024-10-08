package problem363

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test363(t *testing.T) {
	assert.Equal(t, 3, maxSumSubmatrix([][]int{{2, 2, -1}}, 3))
	assert.Equal(t, -1, maxSumSubmatrix([][]int{{2, 2, -1}}, 0))
	assert.Equal(t, -101, maxSumSubmatrix([][]int{{27, 5, -20, -9, 1, 26, 1, 12, 7, -4, 8, 7, -1, 5, 8}, {16, 28, 8, 3, 16, 28, -10, -7, -5, -13, 7, 9, 20, -9, 26}, {24, -14, 20, 23, 25, -16, -15, 8, 8, -6, -14, -6, 12, -19, -13}, {28, 13, -17, 20, -3, -18, 12, 5, 1, 25, 25, -14, 22, 17, 12}, {7, 29, -12, 5, -5, 26, -5, 10, -5, 24, -9, -19, 20, 0, 18}, {-7, -11, -8, 12, 19, 18, -15, 17, 7, -1, -11, -10, -1, 25, 17}, {-3, -20, -20, -7, 14, -12, 22, 1, -9, 11, 14, -16, -5, -12, 14}, {-20, -4, -17, 3, 3, -18, 22, -13, -1, 16, -11, 29, 17, -2, 22}, {23, -15, 24, 26, 28, -13, 10, 18, -6, 29, 27, -19, -19, -8, 0}, {5, 9, 23, 11, -4, -20, 18, 29, -6, -4, -11, 21, -6, 24, 12}, {13, 16, 0, -20, 22, 21, 26, -3, 15, 14, 26, 17, 19, 20, -5}, {15, 1, 22, -6, 1, -9, 0, 21, 12, 27, 5, 8, 8, 18, -1}, {15, 29, 13, 6, -11, 7, -6, 27, 22, 18, 22, -3, -9, 20, 14}, {26, -6, 12, -10, 0, 26, 10, 1, 11, -10, -16, -18, 29, 8, -8}, {-19, 14, 15, 18, -10, 24, -9, -7, -19, -14, 23, 23, 17, -5, 6}}, -100))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	count := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		count[i] = make([]int, len(matrix[i]))

		for j := 0; j < len(count[i]); j++ {
			if i == 0 && j == 0 {
				count[i][j] = matrix[i][j]
			} else if i == 0 {
				count[i][j] = count[i][j-1] + matrix[i][j]
			} else if j == 0 {
				count[i][j] = count[i-1][j] + matrix[i][j]
			} else {
				count[i][j] = matrix[i][j] + count[i-1][j] + count[i][j-1] - count[i-1][j-1]
			}
			if count[i][j] == k {
				return k
			}
		}
	}

	maxArea := math.MinInt
	for i := 0; i < len(count); i++ {
		for j := 0; j < len(count[i]); j++ {
			for x := i; x < len(count); x++ {
				for y := j; y < len(count[x]); y++ {
					area := 0
					if i == x && j == y {
						area = count[x][y]
					} else if i == x {
						area = count[x][y] - count[x][j]
					} else if j == y {
						area = count[x][y] - count[i][y]
					} else {
						area = count[x][y] + count[i][j] - count[x][j] - count[i][y]
					}
					if area == k {
						return area
					}
					if area <= k {
						maxArea = max(maxArea, area)
					}
				}
			}
		}
	}

	return maxArea
}
