package problem1557

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	endSet := make(map[int]struct{})
	for _, edge := range edges {
		endSet[edge[1]] = struct{}{}
	}
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if _, e := endSet[i]; !e {
			result = append(result, i)
		}
	}

	return result
}
