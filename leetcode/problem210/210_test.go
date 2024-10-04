package problem210

func findOrder(numCourses int, prerequisites [][]int) []int {
	coursePreq := make(map[int]map[int]struct{})
	indeg := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		if coursePreq[prerequisite[1]] == nil {
			coursePreq[prerequisite[1]] = map[int]struct{}{}
		}
		coursePreq[prerequisite[1]][prerequisite[0]] = struct{}{}
		indeg[prerequisite[0]]++
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0)
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		result = append(result, first)
		for k, _ := range coursePreq[first] {
			indeg[k]--
			if indeg[k] == 0 {
				queue = append(queue, k)
			}
		}
	}
	if len(result) != numCourses {
		return []int{}
	}
	return result
}
