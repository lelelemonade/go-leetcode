package problem542

func updateMatrix(mat [][]int) [][]int {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	rows := len(mat)
	cols := len(mat[0])

	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			if mat[i][j] == 0 {
				dist[i][j] = 0
			} else {
				dist[i][j] = 1<<31 - 1
			}
		}
	}

	queue := make([][2]int, 0)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]

		for _, dir := range dirs {
			newX, newY := x+dir[0], y+dir[1]

			if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
				if dist[newX][newY] > dist[x][y]+1 {
					dist[newX][newY] = dist[x][y] + 1
					queue = append(queue, [2]int{newX, newY})
				}
			}
		}
	}

	return dist
}
