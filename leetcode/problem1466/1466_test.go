package problem1466

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minReorder(n int, connections [][]int) int {
	alreadyConnect := make(map[int]struct{})
	visited := make([]bool, len(connections))
	alreadyConnect[0] = struct{}{}

	modifyCount := 0

	for len(alreadyConnect) != n {
		for i := 0; i < len(connections); i++ {
			if visited[i] {
				continue
			}
			connection := connections[i]
			_, end := alreadyConnect[connection[0]]
			_, start := alreadyConnect[connection[1]]
			if !start && end {
				alreadyConnect[connection[1]] = struct{}{}
				visited[i] = true
			}
			if start && !end {
				connection[1], connection[0] = connection[0], connection[1]
				alreadyConnect[connection[1]] = struct{}{}
				modifyCount++
				visited[i] = true
			}
		}
	}
	return len(connections) - modifyCount
}

func Test1466(t *testing.T) {
	// assert.Equal(t,3,minReorder(6,[][]int{{0,1},{1,3},{2,3},{4,0},{4,5}}))
	assert.Equal(t, 3, minReorder(6, [][]int{{0, 2}, {0, 3}, {4, 1}, {4, 5}, {5, 0}}))
}
