package problem207

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	graph := make(map[int]map[int]struct{})

	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		if graph[prerequisite] == nil {
			graph[prerequisite] = make(map[int]struct{})
		}
		graph[prerequisite][course] = struct{}{}
		indegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	completedCourses := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		completedCourses++

		for dependentCourse, _ := range graph[course] {
			indegree[dependentCourse]--
			if indegree[dependentCourse] == 0 {
				queue = append(queue, dependentCourse)
			}
		}
	}

	return completedCourses == numCourses
}
