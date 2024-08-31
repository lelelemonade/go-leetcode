package problem2699

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	foundMinPath := make(map[int][]int)
	foundMinPath[source] = []int{0, -1, source}

	for len(foundMinPath) != n {
		minusOneLocation := -1
		minWeight := math.MaxInt
		minStart := 0
		minEnd := 0
		for i := 0; i < len(edges); i++ {
			if foundMinPath[edges[i][0]] != nil &&
				foundMinPath[edges[i][1]] == nil &&
				edges[i][2]+foundMinPath[edges[i][0]][0] < minWeight {
				minStart = edges[i][0]
				minEnd = edges[i][1]
				if edges[i][2] == -1 {
					minusOneLocation = i
					edges[i][2] += 2
				}
				minWeight = edges[i][2] + foundMinPath[edges[i][0]][0]
			} else if foundMinPath[edges[i][1]] != nil &&
				foundMinPath[edges[i][0]] == nil &&
				edges[i][2]+foundMinPath[edges[i][1]][0] < minWeight {
				minStart = edges[i][1]
				minEnd = edges[i][0]
				if edges[i][2] == -1 {
					minusOneLocation = i
					edges[i][2] += 2
				}
				minWeight = edges[i][2] + foundMinPath[edges[i][1]][0]
			}
		}
		foundMinEnd := make([]int, len(foundMinPath[minStart]))
		copy(foundMinEnd, foundMinPath[minStart])
		foundMinEnd[0] = minWeight
		if minusOneLocation != -1 {
			foundMinEnd[1] = minusOneLocation
		}
		foundMinPath[minEnd] = append(foundMinEnd, minEnd)
	}

	minPath := foundMinPath[destination]

	if minPath[0] > target || (minPath[0] < target && minPath[1] == -1) {
		return [][]int{}
	} else {
		edges[minPath[1]][2] += target - minPath[0]
		return edges
	}
}

func Test2699(t *testing.T) {
	assert.Equal(
		t,
		[][]int{
			{1, 0, 4},
			{1, 2, 3},
			{2, 3, 5},
			{0, 3, 1},
		},
		modifiedGraphEdges(
			4,
			[][]int{
				{1, 0, 4},
				{1, 2, 3},
				{2, 3, 5},
				{0, 3, -1},
			},
			0,
			2,
			6,
		),
	)

	// assert.Equal(
	// 	t,
	// 	[][]int{
	// 		{1, 0, 4},
	// 		{1, 2, 3},
	// 		{2, 3, 5},
	// 		{0, 3, 1},
	// 	},
	// 	modifiedGraphEdges(
	// 		4,
	// 		[][]int{
	// 			{1, 0, 4},
	// 			{1, 2, 3},
	// 			{2, 3, 5},
	// 			{0, 3, -1},
	// 		},
	// 		0,
	// 		2,
	// 		6,
	// 	),
	// )

	// assert.Equal(
	// 	t,
	// 	[][]int{
	// 		{4, 1, 1},
	// 		{2, 0, 1},
	// 		{0, 3, 1},
	// 		{4, 3, 3},
	// 	},
	// 	modifiedGraphEdges(
	// 		5,
	// 		[][]int{
	// 			{4, 1, -1},
	// 			{2, 0, -1},
	// 			{0, 3, -1},
	// 			{4, 3, -1},
	// 		},
	// 		0,
	// 		1,
	// 		5,
	// 	),
	// )

	// assert.Equal(
	// 	t,
	// 	[][]int{},
	// 	modifiedGraphEdges(
	// 		3,
	// 		[][]int{
	// 			{0, 1, -1},
	// 			{0, 2, 5},
	// 		},
	// 		0,
	// 		2,
	// 		6,
	// 	),
	// )
}
