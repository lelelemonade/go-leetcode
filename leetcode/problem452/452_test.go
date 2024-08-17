package problem452

import (
	"fmt"
	"sort"
	"testing"
)

func Test452(t *testing.T) {
	//points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	//points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	points := [][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}

	fmt.Println(findMinArrowShots(points))
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	minArrow := 0

	for i := 0; i < len(points); {
		arrowValue := points[i][1]
		minArrow++
		for i < len(points) && arrowValue >= points[i][0] {
			i++
		}
	}

	return minArrow
}
