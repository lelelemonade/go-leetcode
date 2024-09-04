package problem149

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type linearLine struct {
}

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	maxPointCount := 2
	for i := 0; i < len(points); i++ {
		for j := i+1; j < len(points); j++ {
			pointCount := 2
			for k := j+1; k < len(points); k++ {
				if (points[i][0]-points[j][0])*(points[k][1]-points[j][1]) ==
					(points[k][0]-points[j][0])*(points[i][1]-points[j][1]) {
						pointCount++
				}
			}
			maxPointCount=max(maxPointCount,pointCount)
		}
	}
	return maxPointCount
}

func Test149(t *testing.T){
	assert.Equal(t,3,maxPoints([][]int{{1,1},{2,2},{3,3}}))
}
