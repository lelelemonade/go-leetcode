package problem1162

import "math"

type gridNode struct {
	x        int
	y        int
	distance int
}

func maxDistance(grid [][]int) int {
	n := len(grid)
	directions := [][]int{
		// {1, 1},
		// {1, -1},
		// {-1, 1},
		// {-1, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = math.MaxInt
		}
	}
	queue := make([]gridNode, 0)
	landFound := false
	waterFound := false
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, gridNode{i, j, 0})
				dist[i][j] = 0
				landFound = true
			}else{
				waterFound = true
			}
		}
	}
	if !landFound || !waterFound {
		return -1
	}
	for len(queue) != 0 {
		first := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			x := first.x + direction[0]
			y := first.y + direction[1]
			if x < 0 || y < 0 || x >= n || y >= n || grid[x][y] == 1 {
				continue
			}
			if first.distance+1 < dist[x][y] {
				dist[x][y] = first.distance + 1
				queue = append(queue, gridNode{x, y, first.distance + 1})
			}
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && dist[i][j] != math.MaxInt {
				result = max(result, dist[i][j])
			}
		}
	}
	return result
}
