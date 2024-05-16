package problem417

import (
	"fmt"
	"testing"
)

func Test417(t *testing.T) {
	testCase1 := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}

	fmt.Println(pacificAtlantic(testCase1))

	testCase2 := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}

	fmt.Println(pacificAtlantic(testCase2))
}

const flowToPacific int8 = 1
const flowToAtlantic int8 = 2
const allDirection int8 = 3

func pacificAtlantic(heights [][]int) [][]int {
	flow := make([][]int8, len(heights))
	for i, height := range heights {
		flow[i] = make([]int8, len(height))
	}

	var dfs = func(x, y int, flowDirection int8) {}
	dfs = func(x, y int, flowDirection int8) {
		if flow[x][y] == flowDirection ||
			flow[x][y] == allDirection {
			return
		}

		flow[x][y] += flowDirection

		if x+1 < len(heights) && heights[x+1][y] >= heights[x][y] {
			dfs(x+1, y, flowDirection)
		}
		if y+1 < len(heights[x]) && heights[x][y+1] >= heights[x][y] {
			dfs(x, y+1, flowDirection)
		}
		if x-1 >= 0 && heights[x-1][y] >= heights[x][y] {
			dfs(x-1, y, flowDirection)
		}
		if y-1 >= 0 && heights[x][y-1] >= heights[x][y] {
			dfs(x, y-1, flowDirection)
		}
	}

	for i := 0; i < len(heights); i++ {
		dfs(i, 0, flowToPacific)
	}
	for j := 0; j < len(heights[0]); j++ {
		dfs(0, j, flowToPacific)
	}

	for i := len(heights) - 1; i >= 0; i-- {
		dfs(i, len(heights[i])-1, flowToAtlantic)
	}
	for j := len(heights[0]) - 1; j >= 0; j-- {
		dfs(len(heights)-1, j, flowToAtlantic)
	}

	result := make([][]int, 0)

	for i := 0; i < len(flow); i++ {
		for j := 0; j < len(flow[i]); j++ {
			if flow[i][j] == allDirection {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
