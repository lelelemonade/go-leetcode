package problem684

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test684(t *testing.T) {
	//testCase1 := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	//
	//assert.Equal(t, []int{1, 4}, findRedundantConnection(testCase1))

	testCase2 := [][]int{{1, 2}, {1, 3}, {2, 3}}

	assert.Equal(t, []int{2, 3}, findRedundantConnection(testCase2))
}

func findRedundantConnection(edges [][]int) []int {
	unionFind := NewUnionFind(len(edges))
	for i := 0; i < len(edges); i++ {
		if unionFind.union(edges[i][0], edges[i][1]) {
			return edges[i]
		}
	}

	return []int{-1, -1}
}

type UnionFind struct {
	f []int
}

func NewUnionFind(n int) UnionFind {
	f := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		f[i] = i
	}

	return UnionFind{f: f}
}

func (u *UnionFind) union(x, y int) bool {
	xParent := u.find(x)
	yParent := u.find(y)
	if xParent == yParent {
		return true
	}
	u.f[xParent] = yParent
	return false
}

func (u *UnionFind) find(x int) int {
	value := u.f[x]
	if value == x {
		return value
	}
	return u.find(value)
}
