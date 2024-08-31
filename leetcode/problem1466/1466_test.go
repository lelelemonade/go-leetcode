package problem1466

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func dfs(x, parent int, e [][][]int) int {
    res := 0
    for _, edge := range e[x] {
        if edge[0] == parent {
            continue
        }
        res += edge[1] + dfs(edge[0], x, e)
    }
    return res
}

func minReorder(n int, connections [][]int) int {
    e := make([][][]int, n)
    for _, edge := range connections {
        e[edge[0]] = append(e[edge[0]], []int{edge[1], 1})
        e[edge[1]] = append(e[edge[1]], []int{edge[0], 0})
    }
    return dfs(0, -1, e)
}

func Test1466(t *testing.T) {
	assert.Equal(t,3,minReorder(6,[][]int{{0,1},{1,3},{2,3},{4,0},{4,5}}))
	// assert.Equal(t, 3, minReorder(6, [][]int{{0, 2}, {0, 3}, {4, 1}, {4, 5}, {5, 0}}))
}
